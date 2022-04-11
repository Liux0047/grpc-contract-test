package example

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/liux0047/grpc-contract-test/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

var (
	port = flag.Int("port", 50051, "The server port")
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func TestProviderContract(t *testing.T) {
	go func() {
		startServer(t)
	}()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewShoppingCartClient(conn)

	contract, err := contract.ReadConctract("fooshop.textproto")
	if err != nil {
		t.Fatalf("reading contract failed: %v", err)
	}
	for _, interaction := range contract.Interactions {
		switch interaction.Method {
		case "AddItem":
			req := new(pb.AddItemRequest)
			wantResp := new(pb.AddItemResponse)
			interaction.Request.UnmarshalTo(req)
			interaction.Response.UnmarshalTo(wantResp)
			response, err := c.AddItem(context.Background(), req)
			if err != nil {
				t.Fatalf("not able to call AddItem with %v: %v", req, err)
			}
			if !proto.Equal(response, wantResp) {
				t.Errorf("response not conforming to contract")
			}
		}
	}
}

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

func startServer(t *testing.T) {
	t.Helper()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShoppingCartServer(s, &ShoppingCartServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
