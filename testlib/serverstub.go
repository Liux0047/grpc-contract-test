package testlib

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/google/go-cmp/cmp"
	pb "github.com/liux0047/grpc-contract-test/contract"
	// spb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

type ServerStub struct {
	contract *pb.Contract
}

func (stub *ServerStub) Respond(method string, req proto.Message) (proto.Message, error) {
	for _, interaction := range stub.contract.Interactions {
		fmt.Printf("interaction: %v; method equal: %v\n", interaction, interaction.Method == method)
		fmt.Printf("proto diff: %v\n", cmp.Diff(req, interaction.Request, protocmp.Transform()))
		if interaction.Method == method && proto.Equal(req, interaction.Request) {
			return interaction.Response, nil
		}
	}
	return nil, fmt.Errorf("no contract found for %q with request: %v", method, req)
}

func NewServer(filename string) *ServerStub {
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

	// contract := &pb.Contract{}
	// if err := prototext.Unmarshal(content, contract); err != nil {
	// 	log.Fatalln("NewServer: Failed to parse textproto:", err)
	// }
	// fmt.Println(prototext.Format(ct))
	return &ServerStub{parse(filename)}
}

func parse(file string) *pb.Contract {
	// Read the contract defined as a textproto file.
	in, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	contract := &pb.Contract{}
	if err := prototext.Unmarshal(in, contract); err != nil {
		log.Fatalln("Failed to parse textproto:", err)
	}
	return contract
}
