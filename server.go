package main

import (
	pb "github.com/vijayChaurasia-6/user-service/proto"
	"github.com/vijayChaurasia-6/user-service/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50052")
	// Check for errors
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Instantiate the server
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
