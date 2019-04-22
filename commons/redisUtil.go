package commons

import (
	"github.com/go-redis/redis"
	"github.com/lizhixin1992/test/conf"
)

var Cache = New()

func New() *redis.Client {
	conf.GlobalConf.G
	return redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
