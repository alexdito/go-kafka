package queryStorage

import (
	"KafkaProject/model"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
)

var ctx = context.Background()

func Set(balance model.Balance) {
	rdb := GetClient()

	hash := fmt.Sprintf("%x", sha1.Sum([]byte(balance.User)))

	b, err := json.Marshal(balance)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rdb.Set(ctx, hash, b, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, hash).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}

func Get(user *model.User) {
	rdb := GetClient()

	hash := fmt.Sprintf("%x", sha1.Sum([]byte(user.User)))

	val, err := rdb.Get(ctx, hash).Result()

	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}