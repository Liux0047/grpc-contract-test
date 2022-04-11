package contract

import (
	"io/ioutil"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
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
	GotResponse, WantResponse proto.Message
	RpcError, Err             error
}

type RpcTester interface {
	CallRpc(*testing.T, *Interaction) *EvalResult
	RegisterServer(*grpc.Server)
}

func TestProviderContract(t *testing.T, tester RpcTester, addr string) {
	// Start the Shopping cart server as a separate goroutine.
	go func() {
		startServer(t, tester, addr)
	}()

	// Verify conformance to the contract for each interaction.
	contract, err := ReadConctract("fooshop.textproto")
	if err != nil {
		t.Fatalf("reading contract failed: %v", err)
	}
	for _, interaction := range contract.Interactions {
		res := tester.CallRpc(t, interaction)

		if err != nil && !interaction.WantError {
			t.Fatalf("unexpected error in calling %v with %v: %v", interaction.Method, interaction.Request, err)
		}
		if !proto.Equal(res.GotResponse, res.WantResponse) {
			t.Errorf("response not conforming to contract")
		}
	}
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
