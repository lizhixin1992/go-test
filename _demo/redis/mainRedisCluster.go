package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	//集群
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"},
		Password: "",
	})

	defer client.Close()
	fmt.Println(client.Set("test1", "redisUtil", 1*time.Hour).Result())
	fmt.Println(client.Get("test").Result())
}
