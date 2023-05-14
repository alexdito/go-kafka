package eventStorage

import (
	"github.com/redis/go-redis/v9"
)

func GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "event-storage-redis:6379",
		Password: "",
		DB:       0,
	})
}
