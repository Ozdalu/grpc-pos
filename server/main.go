package main

import (
	"context"
	"log"
	"net"

	protobuf "github.com/ozdalu/grpc-go/proto"
	"google.golang.org/grpc"
)

type HelloServiceServer struct {
	protobuf.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) SayHello(ctx context.Context, in *protobuf.InputRequest) (*protobuf.OutputRequest, error) {
	log.Printf("Received: %v", in.GetName())
	return &protobuf.OutputRequest{Message: "Hello " + in.GetName()}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc_server := grpc.NewServer()
	protobuf.RegisterHelloServiceServer(grpc_server, &HelloServiceServer{})
	log.Printf("server listening at %v", listener.Addr())

	if err := grpc_server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
