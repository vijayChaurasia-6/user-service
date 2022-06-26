package main

import (
	"context"
	"fmt"
	pb "github.com/vijayChaurasia-6/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	item, err := c.Save(ctx, &pb.UserRequest{Id: "1", Name: "Vijay"})
	fmt.Println(item.Id)

	if err != nil {
		fmt.Printf("cant connct %v", err)
	}
	item2, err := c.Load(ctx, &pb.UserResponse{Id: "1"})
	fmt.Println(item2.Id)
	if err != nil {
		fmt.Printf("cant connct %v", err)
	}
	//fmt.Println(item2.Id)
}
