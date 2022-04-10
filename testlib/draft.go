package testlib

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/google/go-cmp/cmp"
	pb "github.com/liux0047/grpc-contract-test/contract"
	// spb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

type Draft struct {
	contract *pb.Contract
}

func (stub *Draft) Respond(method string, req proto.Message) (proto.Message, error) {
	for _, interaction := range stub.contract.Interactions {
		fmt.Printf("interaction: %v; method equal: %v\n", interaction, interaction.Method == method)
		fmt.Printf("proto diff: %v\n", cmp.Diff(req, interaction.Request, protocmp.Transform()))
		if interaction.Method == method && proto.Equal(req, interaction.Request) {
			return interaction.Response, nil
		}
	}
	return nil, fmt.Errorf("no contract found for %q with request: %v", method, req)
}

func NewDraft(service, consumer string) *Draft {
	// request, _ := anypb.New(&spb.AddItemRequest{ItemId: 1})
	// resp, _ := anypb.New(&spb.AddItemResponse{Added: true})
	// ct := &pb.Contract{
	// 	Interactions: []*pb.Interaction{
	// 		{
	// 			Request:  request,
	// 			Response: resp,
	// 		},
	// 	},
	// }
	// content, _ := prototext.Marshal(ct)
	// ioutil.WriteFile("example_content.textproto", content, 0)

	//
	// if err := prototext.Unmarshal(content, contract); err != nil {
	// 	log.Fatalln("NewServer: Failed to parse textproto:", err)
	// }
	// fmt.Println(prototext.Format(ct))
	return &Draft{&pb.Contract{
		Service:      service,
		Consumer:     consumer,
		Interactions: make([]*pb.Interaction, 0),
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

	d.contract.Interactions = append(d.contract.Interactions, &pb.Interaction{
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
	return ioutil.WriteFile(d.contract.Consumer+".textproto", content, 0)
}

// Reads the contract defined as a textproto file.
func ReadConctract(file string) (*pb.Contract, error) {
	in, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("Error reading file:", err)
		return nil, err
	}
	contract := &pb.Contract{}
	if err := prototext.Unmarshal(in, contract); err != nil {
		log.Fatalln("Failed to parse textproto:", err)
		return nil, err
	}
	return contract, nil
}
