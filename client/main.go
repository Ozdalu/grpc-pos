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

  /* //// REGISTER
	registerResponse, err := grpc_client.Register(ctx, &protobuf.Empty{})
	if err != nil {
		log.Fatalf("Register Error: %v", err)
	}
  log.Printf("\nUuid : %v\nReputation : %d", registerResponse.GetUuid(), registerResponse.GetReputation())
  uuid := registerResponse.GetUuid() */

  const uuid string = "4827309344320181817"

  //// SUBSCRIBE
	subscribeResponse, err := grpc_client.Subscribe(ctx, &protobuf.SubscribeRequest{Uuid: uuid})
	if err != nil {
		log.Fatalf("Subscribe Error: %v", err)
	}
  log.Printf("\nMessage : %v", subscribeResponse.GetMessage())

  //// GETLASTBLOCK
	getLastBlockResponse, err := grpc_client.GetLastBlock(ctx, &protobuf.Empty{})
	if err != nil {
		log.Fatalf("GetLastBlock Error: %v", err)
	}
  log.Printf("\nBlockInfo :\n\tHash : %v\n\tBlock Number : %d\n\tPrevious Block Hash : %v", getLastBlockResponse.GetBlockHash(), getLastBlockResponse.GetBlockNumber(), getLastBlockResponse.GetPreviousBlockHash())

  //// ADD TRANSACTION
  addTransactionResponse, err := grpc_client.AddTransaction(ctx, &protobuf.Transaction{Sender: uuid, Receiver: "6653957458531705520", Amount: 0, Data: "Askip je bake tous les blocs, marre d'être le seul :("})
	if err != nil {
		log.Fatalf("AddTransaction Error: %v", err)
	}
  log.Printf("\nSuccess sending transaction %v", addTransactionResponse)

  //// BAKE BLOCK
  bakeBlockResponse, err := grpc_client.BakeBlock(ctx, &protobuf.BakeRequest{Uuid: uuid})
	if err != nil {
		log.Fatalf("BakeBlock Error: %v", err)
	}
  log.Printf("\nUuid : %v\nMessage : %v", bakeBlockResponse.GetUuid(), bakeBlockResponse.GetMessage())

  //// CONFIRMBAKE
	confirmBakeResponse, err := grpc_client.ConfirmBake(ctx, &protobuf.ConfirmRequest{Uuid: uuid})
	if err != nil {
		log.Fatalf("ConfirmBlake Error: %v", err)
	}
  log.Printf("Confirm Blake success ! %v", confirmBakeResponse)
}
