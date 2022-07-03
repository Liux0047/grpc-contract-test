package contract

import (
	"fmt"
	"log"
	"net"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/testing/protocmp"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

func VerifyProviderContract(contracts []*Contract, tester RpcTester) map[string][]error {
	// Start the Shopping cart server as a separate goroutine.
	go func() {
		startServer(tester, *addr)
	}()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := tester.RegisterClient(conn)
	consumerErrorsMap := make(map[string][]error)

	// Verify conformance to the contract for each interaction.
	for _, contract := range contracts {
		var errors []error
		for _, interaction := range contract.Interactions {
			if err := setupPrecondition(interaction.Preconditions, contract.Interactions, tester); err != nil {
				errors = append(errors, fmt.Errorf(
					"unable to setup precondition %v: %v", interaction.Preconditions, err))
				continue
			}
			res := callRpc(interaction, client)
			if err != nil && !interaction.WantError {
				errors = append(errors, fmt.Errorf(
					"unexpected error in calling %v with %v: %v", interaction.Method, interaction.Request, err))
				continue
			}
			gotResp, err := anypb.New(res.Response)
			if err != nil {
				errors = append(errors, fmt.Errorf(
					"unexpected error in marshalling response %v to anypb: %v", res.Response, err))
				continue
			}
			opts := append(ignoreFieldsWithRules(interaction.Rules), protocmp.Transform())
			if diff := cmp.Diff(gotResp, interaction.Response, opts...); diff != "" {
				errors = append(errors, fmt.Errorf(
					"response not conforming to contract, diff: %v", diff))
			}
			if interaction.Rules != nil {
				ruleMet, err := checkRules(res.Response, interaction.Rules)
				if err != nil {
					errors = append(errors, fmt.Errorf(
						"error in evaluating rules %v: %v", interaction.Rules, err))
					continue
				}
				if !ruleMet {
					errors = append(errors, fmt.Errorf("rules are not met for %v", interaction.Rules))
				}
			}
		}
		consumerErrorsMap[contract.Consumer] = errors
	}
	return consumerErrorsMap
}

func startServer(tester RpcTester, addr string) {
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
