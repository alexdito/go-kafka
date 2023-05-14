package kafka

import (
	"KafkaProject/model"
	queryStorage "KafkaProject/query-storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func StartKafka() {
	config := kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "messages",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(config)

	for {
		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}

		fmt.Println("Message is: ", string(message.Value))

		var balance model.Balance

		err = json.Unmarshal(message.Value, &balance)
		if err != nil {
			log.Println(err)
		}

		queryStorage.Set(balance)
	}
}
