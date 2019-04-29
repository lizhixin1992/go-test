package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	//单机
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		PoolSize:     5,
		MinIdleConns: 3,
	})
	defer client.Close()
	//fmt.Println(client.Get("redisUtil"))
	//fmt.Println(client.PoolStats())

	//wg := sync.WaitGroup{}
	//wg.Add(10)

	//for i := 0; i < 10; i++ {
	//	go func() {
	//		defer wg.Done()

	for j := 0; j < 100; j++ {
		client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
		client.Get(fmt.Sprintf("name%d", j)).Result()
	}

	fmt.Println("PoolStats:", client.PoolStats())
	//}()
	//}

	//wg.Wait()

}
