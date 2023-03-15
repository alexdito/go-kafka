package main

import (
	appKafka "KafkaProject/kafka"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Kafka consumer")

	go appKafka.StartSendingKafka()

	fmt.Println("Kafka has been started...")

	time.Sleep(10 * time.Minute)
}
