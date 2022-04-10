package example

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"testing"

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
	pb.UnimplementedShoppingCartServer
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
		pb.RegisterShoppingCartServer(s, &stubServer{})
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
	resp, err := c.AddItem(context.Background(), &pb.AddItemRequest{ItemId: 1})
	if err != nil {
		log.Fatalf("could not add item: %v", err)
	}
	log.Printf("Got response: %v", resp)

	if !proto.Equal(resp, &pb.AddItemResponse{Added: true}) {
		t.Errorf("incorrect")
	}
}
