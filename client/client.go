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
	fmt.Println(message.Text)

	FirstHandshake, err := c.GetHeader(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}
	//Wait untill server has sent the second handshake
	for FirstHandshake.Ack != 1 {

	}
	fmt.Printf("Client recieved second handshake from server with Syn flag True and Ack: %d \n", FirstHandshake.Ack)

	message = protos.Message{Text: "Third hardshake sent from Client with Seq", Seq: FirstHandshake.Seq + 1}
	fmt.Println(message.Text, message.Seq)

	ThirdHandshake, err := c.GetHeader(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}

	//Data exhange logic here
	fmt.Printf("Recieved from server, Ack: %d \n", ThirdHandshake.Ack)

	for i := 0; i < 10; i++ {
		dataSimulation, err := c.GetHeader(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling GetHeader(Message): %s", err)
		}
		fmt.Printf("Recieved from server, Ack: %d \n", dataSimulation.Ack)
	}
}
