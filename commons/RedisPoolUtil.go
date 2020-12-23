package commons

import (
	"github.com/gomodule/redigo/redis"
	"github.com/lizhixin1992/go-test/conf"
	"github.com/pelletier/go-toml"
	"time"
)

var CachePool = NewPool()

func NewPool() *redis.Pool {
	tree := conf.GlobalConf.Get("redis").(*toml.Tree)
	return &redis.Pool{
		MaxIdle:     30,
		IdleTimeout: 300 * time.Second,
		Wait:        true,
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial(conf.Tcp, tree.Get("Addr").(string),
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
}
