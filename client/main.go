package main

import (
	"context"
	"log"

	pb "github.com/k197781/go-auth-proto"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAuthClient(conn)
	ctx := context.Background()
	r, err := c.CreateToken(ctx, &pb.CreateTokenRequest{Id: "1", Name: "yuna"})
	if err != nil {
		log.Fatalf("could not create token: %v", err)
	}
	log.Printf("token: %s", r.Token)
}
