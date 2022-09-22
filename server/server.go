package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Grumlebob/GoLangAssignment2HardgRPC/protos"

	"google.golang.org/grpc"
)

type Server struct {
	protos.ChatServiceServer
}

func (s *Server) GetHeader(ctx context.Context, message *protos.Message) (*protos.Message, error) {
	fmt.Printf("recieved: %v", message) //måske noget gøgl her? jo, server skal printe et eller andet.
	var msgServer *protos.Message
	if (message.Seq == 0) && (message.Ack == 0) {
		msgServer = &protos.Message{Text: "Second handshake  sent from Server, with Syn flag True and Ack 1", Ack: 1, Seq: 1}
	} else {
		msgServer = &protos.Message{Text: "Exchanging data", Ack: message.Seq + 1, Seq: message.Seq}
	}
	return msgServer, nil
}

func main() {
	// Create listener tcp on port 9080
	listener, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterChatServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
