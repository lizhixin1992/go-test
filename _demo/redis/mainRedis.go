package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	//单机
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})
	defer client.Close()
	fmt.Println(client.Get("redisUtil"))
	fmt.Println(client.PoolStats())

}
