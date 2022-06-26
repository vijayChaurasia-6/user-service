package service

import (
	"context"
	"github.com/vijayChaurasia-6/user-service/model"
	pb "github.com/vijayChaurasia-6/user-service/proto"
	"log"
)

type Server struct {
	pb.UnimplementedUserServer
}

func (s *Server) Save(context.Context, *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received a note to save")
	model.SaveUser()
	return &pb.UserResponse{Id: "1"}, nil

}

func (s *Server) Load(ctx context.Context, response *pb.UserResponse) (*pb.UserRequest, error) {
	log.Printf("Received a note to save: ")
	model.GetUser()
	return &pb.UserRequest{Id: "1", Name: "Vijay", Age: 27}, nil
}
