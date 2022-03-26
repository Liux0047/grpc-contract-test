package example

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	pb "github.com/liux0047/grpc-contract-test/example/shoppingcart"
)

var (
	port = flag.Int("port", 50051, "The server port")
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type stubServer struct {
	lib *gctlib
	pb.UnimplementedShoppingCartServer
}

func (s *stubServer) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	resp, err := s.lib.respond("ShoppingCart.AddItem", req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddItemResponse), nil
}

func TestClientContract(t *testing.T) {
	// ct lib reads the contracts and store map[rpcMethod]list(req, resp)
	go func() {
		flag.Parse()
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterShoppingCartServer(s, &stubServer{lib: NewGctLib()})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewShoppingCartClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.AddItem(ctx, &pb.AddItemRequest{ItemId: 1})
	if err != nil {
		log.Fatalf("could not add item: %v", err)
	}
	log.Printf("Got response: %v", resp)

	if !proto.Equal(resp, &pb.AddItemResponse{Added: true}) {
		t.Errorf("incorrect")
	}
}

// contract test lib
type gctlib struct {
	contracts map[string][]*interaction
}

type interaction struct {
	req  proto.Message
	resp proto.Message
}

func (lib *gctlib) respond(method string, req proto.Message) (proto.Message, error) {
	for _, interaction := range lib.contracts[method] {
		if proto.Equal(req, interaction.req) {
			return interaction.resp, nil
		}
	}
	return nil, fmt.Errorf("no contract found for %q with request: %v", method, req)
}

func NewGctLib() *gctlib {
	return &gctlib{
		contracts: map[string][]*interaction{
			"ShoppingCart.AddItem": {
				&interaction{
					req:  &pb.AddItemRequest{ItemId: 1},
					resp: &pb.AddItemResponse{Added: true},
				},
			},
		},
	}
}
