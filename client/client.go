package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/Grumlebob/GoLangAssignment2HardgRPC/protos"
)

func main() {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := protos.NewChatServiceClient(conn)

	SendRequest(c)
}

func SendRequest(c protos.ChatServiceClient) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	message := protos.Message{Text: "First handshake from client, with Syn flag True and Seq 0", Ack: 0, Seq: 0}

	response, err := c.GetHeader(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}

	fmt.Printf("Third hardshake from Client, server connection established with Ack 1, Text: %v \n", response)
}
