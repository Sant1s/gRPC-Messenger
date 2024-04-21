package server

import (
	"context"
	pb "github.com/Sant1s/gRPC-1C-KIS/pkg/github.com/Sant1s/messenger"
	"sync"
	"time"
)

type Messenger struct {
	pb.UnimplementedMessengerServer
	mu       sync.Mutex
	Messages map[string][]*pb.Message
}

func (s *Messenger) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.MessageNotification, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	recipient := req.GetRecipient()
	message := req.GetMessage()
	if _, ok := s.Messages[recipient]; !ok {
		s.Messages[recipient] = []*pb.Message{}
	}
	s.Messages[recipient] = append(s.Messages[recipient], message)
	return &pb.MessageNotification{Message: &pb.Message{
		Sender: "server",
		Text:   "ok",
	}}, nil
}

func (s *Messenger) ReceiveMessage(req *pb.ReceiveMessageRequest, stream pb.Messenger_ReceiveMessageServer) error {
	recipient := req.GetRecipient()
	for {
		time.Sleep(1 * time.Second)
		s.mu.Lock()
		messages, ok := s.Messages[recipient]
		if ok && len(messages) > 0 {
			for _, msg := range messages {
				if err := stream.Send(&pb.MessageNotification{Message: msg}); err != nil {
					s.mu.Unlock()
					return err
				}
			}
			s.Messages[recipient] = []*pb.Message{}
		}
		s.mu.Unlock()
	}
}
