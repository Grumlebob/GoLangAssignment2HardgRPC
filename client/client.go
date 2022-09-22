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
	message := protos.Message{Text: "Client sent first handshake, with Syn flag True and Seq 0", Ack: 0, Seq: 0}
	fmt.Println(message.Text)

	FirstHandshake, err := c.GetHeader(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}
	//WAIT FOR SERVER TO SEND SECOND HANDSHAKE
	for FirstHandshake.Ack != 1 {

	}
	fmt.Printf("Client recieved second handshake from server with Syn flag True and Ack: %d \n", FirstHandshake.Ack)

	//Third handshake made explicit, but could be part of data exchange loop.
	message = protos.Message{Text: "Client sent third hardshake, with Ack: 1 and Seq: ", Ack: FirstHandshake.Ack, Seq: FirstHandshake.Seq + 1}
	fmt.Println(message.Text, message.Seq)
	ThirdHandshake, err := c.GetHeader(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetHeader(Message): %s", err)
	}
	fmt.Printf("Client Recieved from server, Ack: %d \n", ThirdHandshake.Ack)

	//Data exhange logic here
	for i := 0; i < 5; i++ {
		message.Seq++
		fmt.Printf("Client Sent to server fictional data with Seq: %d \n", message.Seq)
		dataSimulation, err := c.GetHeader(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling GetHeader(Message): %s", err)
		}
		fmt.Printf("Client recieved from server Ack: %d \n", dataSimulation.Ack)
	}
}
