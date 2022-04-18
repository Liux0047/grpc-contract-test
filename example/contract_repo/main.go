package main

import (
	"log"

	"github.com/liux0047/grpc-contract-test/contract"
	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

func generateConsumerContract() {
	draft := contract.NewContract("shoppingcart", "fooshop")
	draft.AddInteraction(
		"Add item to cart",
		"AddItem2",
		&pb.AddItemRequest{ItemId: 1},
		&pb.AddItemResponse{Added: true},
		false, nil, nil)
	if err := draft.PublishRemote(false); err != nil {
		log.Fatalf("Error in publishing to local: %v", err)
	}
}

func main() {
	generateConsumerContract()
}
