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
	message := protos.Message{Text: "Client sent first handshake, with Syn flag True and Seq 0", Ack: 0, Seq: 0}
	FirstHandshake, err := c.GetHeader(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}
	//Wait untill server has sent the second handshake
	for FirstHandshake.Ack != 1 {

	}
	fmt.Printf("Client recieved second handshake from server with Ack: %d \n", FirstHandshake.Ack)

	message = protos.Message{Text: "Client sent first handshake, with Syn flag True and Seq 0", Ack: 0, Seq: 0}
	ThirdHandshake, err := c.GetHeader(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}
	fmt.Printf("Third hardshake sent from Client with Seq %v, And some additional fictional data \n", ThirdHandshake.Seq)

}
