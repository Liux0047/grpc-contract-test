package example

import (
	"testing"

	"github.com/liux0047/grpc-contract-test/testlib"

	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

func TestConsumerContract(t *testing.T) {
	draft := testlib.NewDraft("shoppingcart", "uniqlo")
	draft.AddInteraction("AddItem", &pb.AddItemRequest{ItemId: 1}, &pb.AddItemResponse{Added: true})
	draft.Commit()
}
