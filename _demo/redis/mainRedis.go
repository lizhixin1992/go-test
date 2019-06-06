package main

import (
	"github.com/go-redis/redis"
	"github.com/lizhixin1992/test/commons"
)

func main() {
	//单机
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()
	//fmt.Println(client.Get("redisUtil"))
	//fmt.Println(client.PoolStats())

	//wg := sync.WaitGroup{}
	//wg.Add(10)
	//
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		defer wg.Done()
	//
	//		for j := 0; j < 100; j++ {
	//			client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
	//			client.Get(fmt.Sprintf("name%d", j)).Result()
	//		}
	//
	//		fmt.Println("PoolStats:", client.PoolStats())
	//	}()
	//}
	//
	//wg.Wait()

	//client.Set("test","test11111",30 * time.Second)
	//fmt.Println("PoolStats:", client.PoolStats())

	//v,err := client.Get("test").Result()
	//fmt.Println(v,err)

	//commons.SetString("test","test11111",30 * time.Second)
	//fmt.Println(commons.GetString("test"))

	//fmt.Println(client.SetNX("test1","test11111",30 * time.Second).Result())
	//fmt.Println(commons.SetNXString("test1","test11111",30 * time.Second))

	//fmt.Println(commons.GetSetString("test","newTest"))

	//fmt.Println(client.StrLen("test"))

	//fmt.Println(client.Append("test2","121212").Result())

	//fmt.Println(client.HSet("hset","test","wer3232wew").Result())
	//commons.HSet("hset","test1","速度速度速度")

	//fmt.Println(client.HSetNX("hset","test2","121h23hh23hh2h32").Result())
	//commons.HSetNX("hset","test3","1111111111")

	//fmt.Println(commons.HGet("hset","test3"))

	//fmt.Println(client.HExists("hset","test").Result())
	//fmt.Println(commons.HExists("hset","test"))

	//fmt.Println(client.HDel("hset","test","test1").Result())
	//commons.HDel("hset","test2","test3")

	//data := make(map[string]interface{})
	//	//data["test"] = "wwwwwww"
	//	//data["test1"] = "11111111"
	//	//data["test2"] = "22222222"
	//	//data["test3"] = "33333333"
	//	////fmt.Println(client.HMSet("hset",data).Result())
	//	//commons.HMSet("hset", data)

	//commons.LPush("testList","22222")
	//commons.RPush("testList","end")

	//fmt.Println(commons.LPop("testList"))
	//fmt.Println(commons.RPop("testList"))

	//commons.LRem("testList",0,"end")

	//commons.SAdd("bbb", "111111")

	//fmt.Println(commons.SMembers("a"))
	//fmt.Println(commons.SInter("a","11"))
	//commons.SInterStore("22","a","11")
	//fmt.Println(commons.SUnion("a", "11"))
	//fmt.Println(commons.ZScore("page_rank", "google.com"))
	commons.RenameNX("12121", "test1")
}
