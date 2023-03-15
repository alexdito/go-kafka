package main

import (
	appKafka "KafkaProject/kafka"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Kafka producer")

	go appKafka.StartKafka()

	fmt.Println("Kafka has been started...")

	time.Sleep(10 * time.Minute)
}
