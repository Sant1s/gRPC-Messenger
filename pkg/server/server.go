package server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/Sant1s/gRPC-1C-KIS/pkg/github.com/Sant1s/messenger"
	"github.com/go-redis/redis/v8"
)

type Messenger struct {
	pb.UnimplementedMessengerServer
	redisClient *redis.Client
}

type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"` // Используем имя пользователя получателя как ключ
	Message   string `json:"message"`
	Read      bool   `json:"read"`
}

func (m *Messenger) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.MessageNotification, error) {
	recipient := req.GetRecipient()
	message := req.GetMessage()
	sender := message.GetSender()

	messageJSON, err := json.Marshal(Message{
		Sender:    sender,
		Recipient: recipient,
		Message:   message.GetText(),
		Read:      false,
	})
	if err != nil {
		return &pb.MessageNotification{Message: &pb.Message{
			Sender: "server",
			Text:   err.Error(),
		}}, err
	}

	err = m.redisClient.LPush(context.TODO(), recipient, messageJSON, 0).Err()
	if err != nil {
		return &pb.MessageNotification{Message: &pb.Message{
			Sender: "server",
			Text:   err.Error(),
		}}, err
	}

	return &pb.MessageNotification{Message: &pb.Message{
		Sender: "server",
		Text:   "ok",
	}}, nil
}

func (m *Messenger) ReceiveMessage(req *pb.ReceiveMessageRequest, stream pb.Messenger_ReceiveMessageServer) error {
	recipient := req.GetRecipient()
	for {
		time.Sleep(1 * time.Second)
		readMessage, err := m.redisClient.LPop(context.Background(), recipient[:len(recipient)-1]).Bytes()
		if err != nil {
			if err == redis.Nil {
				continue
			} else {
				log.Fatalln(err.Error())
			}
		}

		var message Message
		err = json.Unmarshal(readMessage, &message)
		if err != nil {
			continue
		}

		if err := stream.Send(&pb.MessageNotification{Message: &pb.Message{
			Sender: message.Sender,
			Text:   message.Message,
		}}); err != nil {
			return err
		}
	}
}

func NewMessenger() *Messenger {
	return &Messenger{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}
