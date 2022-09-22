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
	protos.GetCurrentTimeServer
}

func (s *Server) GetTime(ctx context.Context, in *protos.GetTimeRequest) (*protos.GetTimeReply, error) {
	fmt.Printf("Received XXX request")
	return &protos.GetTimeReply{Reply: "Your reply here"}, nil
}

func main() {
	// Create listener tcp on port 9080
	listener, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterGetCurrentTimeServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
