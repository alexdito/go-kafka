package main

import (
	eventStorage "KafkaProject/event-storage"
	appKafka "KafkaProject/kafka"
	"KafkaProject/model"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	//Set balance command
	r.POST("/transaction", func(c *gin.Context) {
		var data model.Transaction

		if err := c.ShouldBindQuery(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Timestamp = time.Now().UnixNano()

		b, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		//go

		hash := fmt.Sprintf("%x", sha1.Sum([]byte(data.User)))
		fmt.Println(hash)
		transactionList, balance := eventStorage.Set(hash, b)

		b, err = json.Marshal(balance)
		if err != nil {
			fmt.Println(err)
			return
		}

		appKafka.SendMessage(b)

		fmt.Println(balance)
		fmt.Println(transactionList)
		//withdrawal
		//additional

		c.JSON(http.StatusOK, gin.H{"status": "Set balance success"})
	})

	err := r.Run()

	if err != nil {
		return
	}
}
