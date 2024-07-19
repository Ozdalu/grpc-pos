package main

import (
	"context"
	"log"
	"time"

	protobuf "github.com/ozdalu/grpc-pos/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
  connection, err := grpc.NewClient("35.241.224.46:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	grpc_client := protobuf.NewBlockchainClient(connection)
	ctx, CtxCancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer CtxCancel()

	res, err := grpc_client.Register(ctx, &protobuf.Empty{})
	if err != nil {
		log.Fatalf("Register Error: %v", err)
	}

  log.Printf("\nUuid : %v\nReputation : %d", res.GetUuid(), res.GetReputation())
}
