package main

import (
	"log"
	"net"
	"os"

	pb "github.com/k197781/go-auth-proto"
	"github.com/k197781/go-auth/service"
	"google.golang.org/grpc"
)

func main() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth := service.NewAuthService(os.Getenv("SECRET_KEY"))
	pb.RegisterAuthServer(s, auth)
	log.Printf("gRPC server listening on " + addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
