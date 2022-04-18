package contract

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

// Reads the contract defined as a textproto file.
func ReadConctract(file string) (*Contract, error) {
	in, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("Error reading file:", err)
		return nil, err
	}
	contract := &Contract{}
	if err := prototext.Unmarshal(in, contract); err != nil {
		log.Fatalln("Failed to parse textproto:", err)
		return nil, err
	}
	return contract, nil
}

type RpcTester interface {
	RegisterServer(*grpc.Server)                 // Register the gRPC server.
	RegisterClient(*grpc.ClientConn) interface{} // Register the gRPC client.
}

type EvalResult struct { // The result of invoking the rpc.
	Response      proto.Message // The actual response received.
	RpcError, Err error         // Errors in replaying the rpc.
}

func VerifyProviderContract(t *testing.T, tester RpcTester) {
	// Start the Shopping cart server as a separate goroutine.
	go func() {
		startServer(t, tester, *addr)
	}()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := tester.RegisterClient(conn)

	// Verify conformance to the contract for each interaction.
	dir := "contract_repo/server"
	files, err := ioutil.ReadDir(dir)
	// contract, err := ReadConctract("contract_repo/fooshop.textproto")
	if err != nil {
		t.Fatalf("reading contract in directory %s failed: %v", dir, err)
	}
	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			fileName := filepath.Join(dir, file.Name())
			contract, err := ReadConctract(fileName)
			if err != nil {
				t.Fatalf("error reading contract %s: %v", fileName, err)
			}
			for _, interaction := range contract.Interactions {
				if err := setupPrecondition(t, interaction.Preconditions, contract.Interactions, tester); err != nil {
					t.Fatalf("unable to setup precondition %v: %v", interaction.Preconditions, err)
				}
				res := callRpc(interaction, client)
				if err != nil && !interaction.WantError {
					t.Fatalf("unexpected error in calling %v with %v: %v", interaction.Method, interaction.Request, err)
				}
				gotResp, err := anypb.New(res.Response)
				if err != nil {
					t.Fatalf("unexpected error in marshalling response %v to anypb: %v", res.Response, err)
				}
				opts := append(findFieldsWithRules(interaction.Rules), protocmp.Transform())
				if diff := cmp.Diff(gotResp, interaction.Response, opts...); diff != "" {
					t.Errorf("response not conforming to contract, diff: %v", diff)
				}
				if interaction.Rules != nil {
					ruleMet, err := checkRules(res.Response, interaction.Rules)
					if err != nil {
						t.Fatalf("error in evaluating rules %v: %v", interaction.Rules, err)
					}
					if !ruleMet {
						t.Errorf("rules are not met for %v", interaction.Rules)
					}
				}

			}
		})
	}
}

func findFieldsWithRules(rules *CompositeRule) []cmp.Option {
	var opts []cmp.Option
	if rules == nil {
		return opts
	}
	for _, rule := range rules.IntRules {
		opts = append(opts, cmpopts.IgnoreFields(new(int64), rule.Field))
	}
	for _, rule := range rules.DoubleRules {
		opts = append(opts, cmpopts.IgnoreFields(new(float64), rule.Field))
	}
	for _, rule := range rules.StringRules {
		opts = append(opts, cmpopts.IgnoreFields("", rule.Field))
	}
	for _, rule := range rules.NestedRules {
		opts = append(opts, findFieldsWithRules(rule)...)
	}
	return opts
}

func callRpc(interaction *Interaction, client interface{}) *EvalResult {
	method := reflect.ValueOf(client).MethodByName(interaction.Method)
	if method == (reflect.Value{}) {
		return &EvalResult{
			Response: nil,
			RpcError: nil,
			Err:      fmt.Errorf("unknown method %v", interaction.Method),
		}
	}
	// Obtain a zero value for the rpc method's 2nd parameter, the request.
	// The var req has the correct type needed to invoke the rpc from the client.
	req := reflect.New(method.Type().In(1).Elem()).Interface()
	// Unmarshal the request specified in the contract to this new typed request.
	interaction.Request.UnmarshalTo(req.(proto.Message))
	// Invokes the rpc method with default context and provided request.
	ctx := reflect.ValueOf(context.Background())
	result := method.Call([]reflect.Value{ctx, reflect.ValueOf(req)})
	// Convert the rpc error response if not nil.
	errResponse := result[1].Interface()
	var rpcError error
	if errResponse != nil {
		rpcError = errResponse.(error)
	}
	return &EvalResult{
		Response: result[0].Interface().(proto.Message),
		RpcError: rpcError,
		Err:      nil,
	}
}

func setupPrecondition(t *testing.T, preconditions []string, interactions []*Interaction, client interface{}) error {
	for _, precond := range preconditions {
		found := false
		for _, interaction := range interactions {
			if interaction.Name == precond {
				callRpc(interaction, client)
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("precondition %q not found", precond)
		}
	}
	return nil
}

func checkRules(response proto.Message, rules *CompositeRule) (bool, error) {
	for _, rule := range rules.IntRules {
		met, err := checkIntRule(response, rule)
		if err != nil {
			return false, err
		}
		if !met && rules.Operator == CompositeRule_AND {
			return false, nil
		} else if met && rules.Operator == CompositeRule_OR {
			return true, nil
		}
	}
	for _, rule := range rules.DoubleRules {
		met, err := checkDoubleRule(response, rule)
		if err != nil {
			return false, err
		}
		if !met && rules.Operator == CompositeRule_AND {
			return false, nil
		} else if met && rules.Operator == CompositeRule_OR {
			return true, nil
		}
	}
	for _, rule := range rules.StringRules {
		met, err := checkStringRule(response, rule)
		if err != nil {
			return false, err
		}
		if !met && rules.Operator == CompositeRule_AND {
			return false, nil
		} else if met && rules.Operator == CompositeRule_OR {
			return true, nil
		}
	}
	for _, rule := range rules.NestedRules {
		met, err := checkRules(response, rule)
		if err != nil {
			return false, err
		}
		if !met && rules.Operator == CompositeRule_AND {
			return false, nil
		} else if met && rules.Operator == CompositeRule_OR {
			return true, nil
		}
	}
	return rules.Operator == CompositeRule_AND, nil
}

func startServer(t *testing.T, tester RpcTester, addr string) {
	t.Helper()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	tester.RegisterServer(server)
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func fieldValue(m proto.Message, field string) (reflect.Value, error) {
	v := reflect.ValueOf(m)
	for _, name := range strings.Split(field, ".") {
		if v.Kind() != reflect.Struct {
			return reflect.Value{}, fmt.Errorf("invalid field expression %q for %v", field, m)
		}
		v = v.Elem().FieldByName(name)
	}
	return v, nil
}

func checkIntRule(m proto.Message, rule *IntRule) (bool, error) {
	v, err := fieldValue(m, rule.Field)
	if err != nil || !v.CanInt() {
		return false, fmt.Errorf("error in parsing %q as Int: %v", rule.Field, err)
	}
	intVal := v.Int()
	if rule.CheckMin && intVal < rule.Min {
		return false, nil
	}
	if rule.CheckMax && intVal > rule.Max {
		return false, nil
	}
	if len(rule.Allowed) > 0 && containsInt(rule.Allowed, intVal) {
		return false, nil
	}
	return true, nil
}

func checkDoubleRule(m proto.Message, rule *DoubleRule) (bool, error) {
	v, err := fieldValue(m, rule.Field)
	if err != nil || !v.CanFloat() {
		return false, fmt.Errorf("error in parsing %q as float: %v", rule.Field, err)
	}
	floatVal := v.Float()
	if rule.CheckMin && floatVal < rule.Min {
		return false, nil
	}
	if rule.CheckMax && floatVal > rule.Max {
		return false, nil
	}
	return true, nil
}

func checkStringRule(m proto.Message, rule *StringRule) (bool, error) {
	v, err := fieldValue(m, rule.Field)
	if err != nil || v.Kind() != reflect.String {
		return false, fmt.Errorf("error in parsing %q as string: %v", rule.Field, err)
	}
	stringVal := v.String()
	if len(rule.Allowed) > 0 && containsString(rule.Allowed, stringVal) {
		return false, nil
	}
	if rule.MatchRegex != "" {
		regex, err := regexp.Compile(rule.MatchRegex)
		if err != nil {
			return false, fmt.Errorf("invalid regex %q defined at %v: %v", rule.MatchRegex, rule, err)
		}
		return regex.MatchString(stringVal), nil
	}
	return true, nil
}

func containsInt(arr []int64, target int64) bool {
	for _, e := range arr {
		if target == e {
			return true
		}
	}
	return false
}

func containsString(arr []string, target string) bool {
	for _, e := range arr {
		if target == e {
			return true
		}
	}
	return false
}
