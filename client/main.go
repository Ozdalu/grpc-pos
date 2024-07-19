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

  // uuid := register(grpc_client, ctx)
  const uuid string = "4230130246536364507"

  subscribe(uuid, grpc_client, ctx)
  // getLastBlock(grpc_client, ctx)
  addTransaction(uuid, "7354821206227837404", 1, "on vit dans une #société........... j'espère ne pas me faire #slash cette fois.......", grpc_client, ctx)
  // bakeBlock(uuid, grpc_client, ctx)
  // confirmBake(uuid, grpc_client, ctx)
}

func register(grpc_client protobuf.BlockchainClient, ctx context.Context) string {
	registerResponse, err := grpc_client.Register(ctx, &protobuf.Empty{})
	if err != nil {
		log.Fatalf("Register Error: %v", err)
	}

  log.Printf("\nUuid : %v\nReputation : %d", registerResponse.GetUuid(), registerResponse.GetReputation())
  return registerResponse.GetUuid()
}

func subscribe(uuid string, grpc_client protobuf.BlockchainClient, ctx context.Context) string {
	subscribeResponse, err := grpc_client.Subscribe(ctx, &protobuf.SubscribeRequest{Uuid: uuid})
	if err != nil {
		log.Fatalf("Subscribe Error: %v", err)
	}
  log.Printf("\nMessage : %v", subscribeResponse.GetMessage())

  return subscribeResponse.GetMessage()
}

func getLastBlock(grpc_client protobuf.BlockchainClient, ctx context.Context) protobuf.BlockInfo {
	getLastBlockResponse, err := grpc_client.GetLastBlock(ctx, &protobuf.Empty{})
	if err != nil {
		log.Fatalf("GetLastBlock Error: %v", err)
	}
  log.Printf("\nBlockInfo :\n\tHash : %v\n\tBlock Number : %d\n\tPrevious Block Hash : %v", getLastBlockResponse.GetBlockHash(), getLastBlockResponse.GetBlockNumber(), getLastBlockResponse.GetPreviousBlockHash())

  return protobuf.BlockInfo{}
}

func addTransaction(sender string, receiver string, amount int32, data string, grpc_client protobuf.BlockchainClient, ctx context.Context) {
  _, err := grpc_client.AddTransaction(ctx, &protobuf.Transaction{Sender: sender, Receiver: receiver, Amount: amount, Data: data })
	if err != nil {
		log.Fatalf("AddTransaction Error: %v", err)
	}

  log.Printf("\nSuccess sending transaction")
}

func bakeBlock(uuid string, grpc_client protobuf.BlockchainClient, ctx context.Context) protobuf.BakeResponse {
  bakeBlockResponse, err := grpc_client.BakeBlock(ctx, &protobuf.BakeRequest{Uuid: uuid})
	if err != nil {
		log.Fatalf("BakeBlock Error: %v", err)
	}
  log.Printf("\nUuid : %v\nMessage : %v", bakeBlockResponse.GetUuid(), bakeBlockResponse.GetMessage())

  return protobuf.BakeResponse{}
}

func confirmBake(uuid string, grpc_client protobuf.BlockchainClient, ctx context.Context) {
	_, err := grpc_client.ConfirmBake(ctx, &protobuf.ConfirmRequest{Uuid: uuid})
	if err != nil {
		log.Fatalf("ConfirmBake Error: %v", err)
	}

  log.Printf("Confirm Bake success !")
}
