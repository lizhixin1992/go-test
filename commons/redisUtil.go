package commons

import (
	"github.com/go-redis/redis"
)

var Cache = New()

func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
