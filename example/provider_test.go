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
	contract.VerifyProviderContract(t, new(ProviderTester))
}

type ProviderTester struct{}

func (pt *ProviderTester) RegisterClient(conn *grpc.ClientConn) interface{} {
	return pb.NewShoppingCartClient(conn)
}

func (pt *ProviderTester) RegisterServer(s *grpc.Server) {
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
			Added: true,
		}, nil
	}
	return nil, fmt.Errorf("Unkown item")
}
