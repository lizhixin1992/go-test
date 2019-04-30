package commons

import (
	"github.com/go-redis/redis"
	"github.com/lizhixin1992/test/conf"
	"github.com/pelletier/go-toml"
	"log"
	"time"
)

var redisClient = New()

func New() *redis.Client {
	tree := conf.GlobalConf.Get("redis").(*toml.Tree)
	return redis.NewClient(&redis.Options{
		Addr:     tree.Get("Addr").(string),
		Password: "",
		DB:       int(tree.Get("DB").(int64)),
	})
}

func SetString(key, value string, time time.Duration) {
	err := redisClient.Set(key, value, time).Err()
	if err != nil {
		log.Fatal("redis SetString is err, err : ", err)
		//panic(err)
	}
}

//只在键 key 不存在的情况下， 将键 key 的值设置为 value
//若键 key 已经存在， 则 SETNX 命令不做任何动作
func SetNXString(key, value string, time time.Duration) (flag bool) {
	flag, err := redisClient.SetNX(key, value, time).Result()
	if err != nil {
		flag = false
		log.Fatal("redis SetNXString is err, err : ", err)
		//panic(err)
	}
	return flag
}

//和SetNX相反，只在键已经存在时， 才对键进行设置操作
func SetXXString(key, value string, time time.Duration) (flag bool) {
	flag, err := redisClient.SetXX(key, value, time).Result()
	if err != nil {
		flag = false
		log.Fatal("redis SetXXString is err, err : ", err)
		//panic(err)
	}
	return flag
}

func GetString(key string) (value string) {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		log.Fatal("redis GetString is err, err : ", err)
		//panic(err)
	}
	return value
}

//将键 key 的值设为 value ， 并返回键 key 在被设置之前的旧值
//如果键 key 没有旧值， 也即是说， 键 key 在被设置之前并不存在， 那么命令返回 nil,并且按照传入的key/value保存到redis中
func GetSetString(kev, value string) (oldValue string) {
	oldValue, err := redisClient.GetSet(kev, value).Result()
	if err != nil {
		log.Fatal("redis GetString is err, err : ", err)
	}
	return oldValue
}

//如果键 key 已经存在并且它的值是一个字符串， APPEND 命令将把 value 追加到键 key 现有值的末尾
//如果 key 不存在， APPEND 就简单地将键 key 的值设为 value ， 就像执行 SET key value 一样
func AppendString(key, value string) {
	_, err := redisClient.Append(key, value).Result()
	if err != nil {
		log.Fatal("redis AppendString is err, err : ", err)
	}
}

//将哈希表 hash 中域 field 的值设置为 value
func HSet(key, field string, value interface{}) {
	_, err := redisClient.HSet(key, field, value).Result()
	if err != nil {
		log.Fatal("redis HSet is err, err : ", err)
	}
}

//当且仅当域 field 尚未存在于哈希表的情况下， 将它的值设置为 value
//如果给定域已经存在于哈希表当中， 那么命令将放弃执行设置操作
//如果哈希表 hash 不存在， 那么一个新的哈希表将被创建并执行 HSETNX 命令
func HSetNX(key, field string, value interface{}) {
	_, err := redisClient.HSetNX(key, field, value).Result()
	if err != nil {
		log.Fatal("redis HSetNX is err, err : ", err)
	}
}

//返回哈希表中给定域的值
func HGet(key, field string) (value string) {
	value, err := redisClient.HGet(key, field).Result()
	if err != nil {
		log.Fatal("redis HGet is err, err : ", err)
	}
	return value
}

//检查给定域 field 是否存在于哈希表 hash 当中
func HExists(key, field string) (flag bool) {
	flag, err := redisClient.HExists(key, field).Result()
	if err != nil {
		log.Fatal("redis HExists is err, err : ", err)
	}
	return flag
}

//删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
func HDel(key string, fields ...string) {
	for _, value := range fields {
		_, err := redisClient.HDel(key, value).Result()
		if err != nil {
			log.Fatal("redis HDel is err, err : ", err)
		}
	}
}

//返回哈希表 key 中域的数量
func HLen(key string) (i int) {
	size, err := redisClient.HLen(key).Result()
	if err != nil {
		log.Fatal("redis HLen is err, err : ", err)
	}
	return int(size)
}

//同时将多个 field-value (域-值)对设置到哈希表 key 中
//此命令会覆盖哈希表中已存在的域
//如果 key 不存在，一个空哈希表被创建并执行 HMSET 操作
func HMSet(key string, fields map[string]interface{}) {
	_, err := redisClient.HMSet(key, fields).Result()
	if err != nil {
		log.Fatal("redis HMSet is err, err : ", err)
	}
}
