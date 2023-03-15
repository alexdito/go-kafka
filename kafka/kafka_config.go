package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func StartKafka() {
	config := kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "mytopic",
		GroupID:  "g1",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(config)

	for {
		m, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Some error occured", err)
			continue
		}

		fmt.Println("Message is: ", string(m.Value))
	}
}
