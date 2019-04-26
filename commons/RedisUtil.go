package commons

import (
	"github.com/go-redis/redis"
	"github.com/lizhixin1992/test/conf"
	"github.com/pelletier/go-toml"
)

var Cache = New()

func New() *redis.Client {
	tree := conf.GlobalConf.Get("redis").(*toml.Tree)
	return redis.NewClient(&redis.Options{
		Addr:     tree.Get("Addr").(string),
		Password: "",
		DB:       int(tree.Get("DB").(int64)),
	})
}

var CacheCluster = NewCluster()

func NewCluster() *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:7003", "127.0.0.1:7004", "127.0.0.1:7005"},
		Password: "",
	})
}
