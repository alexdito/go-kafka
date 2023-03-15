package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func StartSendingKafka() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "mytopic",
		Balancer: &kafka.LeastBytes{},
	}

	SendMessage(w)

	if err := w.Close(); err != nil {
		fmt.Println("failed to close writer:", err)
	}
}

func SendMessage(w *kafka.Writer) {
	i := 0
	for {
		i++
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte(fmt.Sprintf("Message number: %d", i)),
			})

		if err != nil {
			fmt.Println("failed to write messages:", err)
		}

		time.Sleep(1 * time.Second)
	}
}
