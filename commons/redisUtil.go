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
