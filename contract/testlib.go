package contract

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"
)

type Draft struct {
	contract *Contract
}

func NewContract(service, consumer string) *Draft {
	return &Draft{&Contract{
		Service:      service,
		Consumer:     consumer,
		Interactions: make([]*Interaction, 0),
	}}
}

func (d *Draft) AddInteraction(method string, req proto.Message, response proto.Message) error {
	reqpb, err := anypb.New(req)
	if err != nil {
		return err
	}
	responsepb, err := anypb.New(response)
	if err != nil {
		return err
	}
	d.contract.Interactions = append(d.contract.Interactions, &Interaction{
		Method:   method,
		Request:  reqpb,
		Response: responsepb,
	})
	return nil
}

func (d *Draft) Commit() error {
	content, err := prototext.MarshalOptions{Multiline: true}.Marshal(d.contract)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(d.contract.Consumer+".textproto", content, 0777)
}

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

type EvalResult struct {
	Response      proto.Message
	RpcError, Err error
}

type RpcTester interface {
	CallRpc(*Interaction) *EvalResult
	RegisterServer(*grpc.Server)
	ContractUrl() string
}

func VerifyProviderContract(t *testing.T, tester RpcTester, addr string) {
	// Start the Shopping cart server as a separate goroutine.
	go func() {
		startServer(t, tester, addr)
	}()

	// Verify conformance to the contract for each interaction.
	// TODO: read all files under URL
	contract, err := ReadConctract(tester.ContractUrl())
	if err != nil {
		t.Fatalf("reading contract failed: %v", err)
	}
	for _, interaction := range contract.Interactions {
		t.Run(interaction.Name, func(t *testing.T) {
			if err := setupPrecondition(t, interaction.Preconditions, contract.Interactions, tester); err != nil {
				t.Fatalf("unable to setup precondition %v: %v", interaction.Preconditions, err)
			}
			res := tester.CallRpc(interaction)
			if err != nil && !interaction.WantError {
				t.Fatalf("unexpected error in calling %v with %v: %v", interaction.Method, interaction.Request, err)
			}
			gotResp, err := anypb.New(res.Response)
			if err != nil {
				t.Fatalf("unexpected error in marshalling response %v to anypb: %v", res.Response, err)
			}
			if diff := cmp.Diff(gotResp, interaction.Response, protocmp.Transform()); diff != "" {
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
		})

	}
}

func setupPrecondition(t *testing.T, preconditions []string, interactions []*Interaction, tester RpcTester) error {
	for _, precond := range preconditions {
		found := false
		for _, interaction := range interactions {
			if interaction.Name == precond {
				tester.CallRpc(interaction)
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

func checkRules(response proto.Message, rules *CompositeRules) (bool, error) {
	for _, rule := range rules.IntRules {
		met, err := checkIntRule(response, rule)
		if err != nil {
			return false, err
		}
		if !met && rules.Operator == CompositeRules_AND {
			return false, nil
		} else if met && rules.Operator == CompositeRules_OR {
			return true, nil
		}
	}
	for _, rule := range rules.StringRules {
		met, err := checkStringRule(response, rule)
		if err != nil {
			return false, err
		}
		if !met && rules.Operator == CompositeRules_AND {
			return false, nil
		} else if met && rules.Operator == CompositeRules_OR {
			return true, nil
		}
	}
	for _, rule := range rules.NestedRules {
		met, err := checkRules(response, rule)
		if err != nil {
			return false, err
		}
		if !met && rules.Operator == CompositeRules_AND {
			return false, nil
		} else if met && rules.Operator == CompositeRules_OR {
			return true, nil
		}
	}
	return rules.Operator == CompositeRules_AND, nil
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
	if rule.ChechMax && intVal > rule.Max {
		return false, nil
	}
	if len(rule.Allowed) > 0 && containsInt(rule.Allowed, intVal) {
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
