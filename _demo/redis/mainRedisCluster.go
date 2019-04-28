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
		PoolSize: 30,
	})

	defer client.Close()
	//fmt.Println(client.Set("test1", "redisUtil", 1*time.Hour).Result())
	//fmt.Println(client.Get("test").Result())

	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("key:%d", i)
		v := k
		val, err := client.Set(k, v, 60*time.Second).Result()
		if err != nil {
			panic(err)
		}

		val, err = client.Get(k).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key:", val)
		fmt.Println("pool state final state:", client.PoolStats()) //获取客户端连接池相关信息
	}

	go func() {
		for i := 1000; i < 2000; i++ {
			k := fmt.Sprintf("key:%d", i)
			v := k
			val, err := client.Set(k, v, 60*time.Second).Result()
			if err != nil {
				panic(err)
			}

			val, err = client.Get(k).Result()
			if err != nil {
				panic(err)
			}
			fmt.Println("key:", val)
		}
	}()

	fmt.Println("pool state final state:", client.PoolStats()) //获取客户端连接池相关信息
}
