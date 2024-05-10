package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Определим структуру сообщения
type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
}

func main() {
	// ctx := context.Background()
	// Подключимся к Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Оставьте этот адрес или укажите свой
	})

	// Отправим сообщение
	sender := "user1"
	recipient := "user2"
	messageText := "Привет! Как дела?"

	message := Message{
		Sender:    sender,
		Recipient: recipient,
		Message:   messageText,
		Read:      false,
	}

	// Сериализуем сообщение в JSON
	messageJSON, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	// Сохраним сообщение в Redis
	err = client.LPush(context.TODO(), "mikiele", messageJSON, 0).Err()
	if err != nil {
		panic(err)
	}
	err = client.LPush(context.TODO(), "mikiele", messageJSON, 0).Err()
	if err != nil {
		panic(err)
	}
	err = client.LPush(context.TODO(), "mikiele", messageJSON, 0).Err()
	if err != nil {
		panic(err)
	}
	err = client.LPush(context.TODO(), "mikiele", messageJSON, 0).Err()
	if err != nil {
		panic(err)
	}
	err = client.LPush(context.TODO(), "mikiele", messageJSON, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Сообщение успешно отправлено!")
	// values, err := client.LRange(ctx, "mikiele", 0, -1).Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, value := range values {
	// 	fmt.Println(value)
	// }

	//	for {
	//		messageJSON, err := client.LPop(ctx, "mikiele").Result()
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println(messageJSON)
	//	}
}
