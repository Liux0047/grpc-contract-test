package example

import (
	"context"
	"fmt"
	"testing"

	"github.com/liux0047/grpc-contract-test/contract"
	"google.golang.org/grpc"

	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

func TestProviderContract(t *testing.T) {
	contract.VerifyProviderContract(t, new(ShoppingCartTester))
}

type ShoppingCartTester struct{}

func (pt *ShoppingCartTester) RegisterClient(conn *grpc.ClientConn) interface{} {
	return pb.NewShoppingCartClient(conn)
}

func (pt *ShoppingCartTester) RegisterServer(s *grpc.Server) {
	pb.RegisterShoppingCartServer(s, NewServer())
}

// A very simple server
type ShoppingCartServer struct {
	pb.UnimplementedShoppingCartServer
}

func NewServer() *ShoppingCartServer {
	return &ShoppingCartServer{}
}

func (s *ShoppingCartServer) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	if req.ItemId == 1 {
		return &pb.AddItemResponse{
			Added: false,
		}, nil
	}
	return nil, fmt.Errorf("Unkown item")
}

func TestConsumerAsServer(t *testing.T) {
	conn, err := contract.NewContractConn("contract_repo/server/fooshop.prototxt")
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewShoppingCartClient(conn)
	ctx := context.Background()
	resp, err := client.AddItem(ctx, &pb.AddItemRequest{ItemId: 1})
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Added {
		t.Errorf("Want Added to be true, got %v", resp.Added)
	}
}
