package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func SendMessage(message []byte) {
	w := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "messages",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: message,
		})

	if err != nil {
		fmt.Println("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		fmt.Println("failed to close writer:", err)
	}
}
