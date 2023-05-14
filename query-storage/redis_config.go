package queryStorage

import (
	"github.com/redis/go-redis/v9"
)

func GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "query-storage-redis:6380",
		Password: "",
		DB:       0,
	})
}
