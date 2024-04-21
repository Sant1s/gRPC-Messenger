package client

import (
	"context"
	pb "github.com/Sant1s/gRPC-1C-KIS/pkg/github.com/Sant1s/messenger"
	"log"
)

func SendMessage(client pb.MessengerClient, recipient, sender, text string) {
	message := &pb.Message{
		Sender: sender,
		Text:   text,
	}
	request := &pb.SendMessageRequest{
		Recipient: recipient,
		Message:   message,
	}
	if _, err := client.SendMessage(context.Background(), request); err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
}

func ReceiveMessages(ctx context.Context, client pb.MessengerClient, recipient string) {
	request := &pb.ReceiveMessageRequest{
		Recipient: recipient,
	}
	stream, err := client.ReceiveMessage(context.Background(), request)
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		log.Printf("Message from %s: %s", msg.Message.Sender, msg.Message.Text)
	}
}
