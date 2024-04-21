package main

import (
	"bufio"
	"context"
	client "github.com/Sant1s/gRPC-1C-KIS/client"
	pb "github.com/Sant1s/gRPC-1C-KIS/pkg/github.com/Sant1s/messenger"
	"google.golang.org/grpc"
	"log"
	"os"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	cli := pb.NewMessengerClient(conn)

	log.Println("Enter recipient:")
	scanner.Scan()
	recipient := scanner.Text()
	recipient = recipient[:len(recipient)-1]

	log.Println("Enter your nickname:")
	scanner.Scan()
	sender := scanner.Text()

	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go client.ReceiveMessages(ctx, cli, recipient)

	for {
		log.Println("Enter message (or 'exit' to quit):")
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		client.SendMessage(cli, recipient, sender, text)
	}
	cancel()
	log.Println("Goodbye!")
}
