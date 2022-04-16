package main

import (
	"github.com/liux0047/grpc-contract-test/contract"
	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

func generateConsumerContract() {
	draft := contract.NewContract("shoppingcart", "fooshop")
	draft.AddInteraction(
		"Add item to cart",
		"AddItem",
		&pb.AddItemRequest{ItemId: 1},
		&pb.AddItemResponse{Added: true},
		false, nil, nil)
	draft.Publish(false)
}

func main() {
	generateConsumerContract()
}
