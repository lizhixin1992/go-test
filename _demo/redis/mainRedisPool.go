package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {
	redisPool := redis.Pool{
		MaxIdle:     30,
		IdleTimeout: 300 * time.Second,
		Wait:        true,
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial("tcp", "127.0.0.1:6379",
				redis.DialPassword(""),
				redis.DialDatabase(0),
				redis.DialConnectTimeout(5*time.Second),
				redis.DialReadTimeout(3*time.Second),
				redis.DialWriteTimeout(3*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	rc := redisPool.Get()
	defer rc.Close()

	v, err := redis.String(rc.Do("GET", "redisUtil"))
	fmt.Println(v)
	fmt.Println(err)
}
