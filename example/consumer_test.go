package example

import (
	"testing"

	"github.com/liux0047/grpc-contract-test/contract"
	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

func TestConsumerContract(t *testing.T) {
	draft := contract.NewContract("shoppingcart", "fooshop")
	draft.AddInteraction("AddItem", &pb.AddItemRequest{ItemId: 1}, &pb.AddItemResponse{Added: true})
	draft.Commit()
}
