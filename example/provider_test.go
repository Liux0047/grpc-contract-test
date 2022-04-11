package example

import (
	"context"
	"flag"
	"fmt"
	"log"
	"testing"

	"github.com/liux0047/grpc-contract-test/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func TestProviderContract(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewShoppingCartClient(conn)

	pt := &ProviderTester{
		client: c,
		server: &ShoppingCartServer{},
	}
	contract.VerifyProviderContract(t, pt, *addr)
}

type ProviderTester struct {
	client pb.ShoppingCartClient
	server pb.ShoppingCartServer
}

func (pt *ProviderTester) CallRpc(t *testing.T, interaction *contract.Interaction) *contract.EvalResult {
	switch interaction.Method {
	case "AddItem":
		req := new(pb.AddItemRequest)
		wantResp := new(pb.AddItemResponse)
		interaction.Request.UnmarshalTo(req)
		interaction.Response.UnmarshalTo(wantResp)
		gotResp, rpcErr := pt.client.AddItem(context.Background(), req)
		return &contract.EvalResult{
			GotResponse:  gotResp,
			WantResponse: wantResp,
			RpcError:     rpcErr,
			Err:          nil,
		}
	default:
		return &contract.EvalResult{
			GotResponse:  nil,
			WantResponse: nil,
			RpcError:     nil,
			Err:          fmt.Errorf("unknown method %v", interaction.Method),
		}
	}
}

func (pt *ProviderTester) RegisterServer(s *grpc.Server) {
	pb.RegisterShoppingCartServer(s, pt.server)
}

// A very simple server
type ShoppingCartServer struct {
	pb.UnimplementedShoppingCartServer
}

func (s *ShoppingCartServer) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	if req.ItemId == 1 {
		return &pb.AddItemResponse{
			Added: true,
		}, nil
	}
	return nil, fmt.Errorf("Unkown item")
}
