package main

import (
	"KafkaProject/model"
	queryStorage "KafkaProject/query-storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//Set balance command
	r.POST("/balance", func(c *gin.Context) {
		var user model.User

		if err := c.ShouldBindQuery(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		queryStorage.Get(&user)

		fmt.Println(user)

		c.JSON(http.StatusOK, gin.H{"status": "success", "user": user.User, "balance": user.Amount})
	})

	err := r.Run(":8081")

	if err != nil {
		return
	}
}
