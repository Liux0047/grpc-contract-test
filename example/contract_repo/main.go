package main

import (
	"github.com/liux0047/grpc-contract-test/contract"
	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

func generateConsumerContract() {
	draft := contract.NewContract("shoppingcart", "fooshop")
	draft.AddInteraction("AddItem", &pb.AddItemRequest{ItemId: 1}, &pb.AddItemResponse{Added: true})
	draft.AddInteraction("AddItem2", &pb.AddItemRequest{ItemId: 1}, &pb.AddItemResponse{Added: true})
	draft.Publish(false)
}

func main() {
	generateConsumerContract()
}
