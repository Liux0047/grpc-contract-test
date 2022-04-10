package example

import (
	"testing"

	"github.com/liux0047/grpc-contract-test/testlib"

	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

func TestProviderContract(t *testing.T) {
	contract, err := testlib.ReadConctract("uniqlo.textproto")
	if err != nil {
		t.Fatalf("reading contract failed: %v", err)
	}
	for _, interaction := range contract.Interactions {
		req := new(pb.AddItemRequest)
		interaction.Request.UnmarshalTo(req)
	}
}
