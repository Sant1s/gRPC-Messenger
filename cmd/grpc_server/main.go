package main

import (
	pb "github.com/Sant1s/gRPC-1C-KIS/pkg/github.com/Sant1s/messenger"
	"github.com/Sant1s/gRPC-1C-KIS/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterMessengerServer(srv, &server.Messenger{
		Messages: make(map[string][]*pb.Message),
	})
	log.Println("Server listening on port 50051")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
