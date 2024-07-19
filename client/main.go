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
  log.Printf("Hello World")
	/* connection, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpc_client := protobuf.NewHelloServiceClient(connection)
	ctx, CtxCancel := context.WithTimeout(context.Background(), time.Second)

	r, err := grpc_client.SayHello(ctx, &protobuf.InputRequest{Name: "Toto"})
	if err != nil {
		log.Fatalf("SayHello Method Error: %v", err)
	}

  log.Printf("Bien reçu : %s", r.GetMessage())

	defer connection.Close()
	defer CtxCancel() */
}
