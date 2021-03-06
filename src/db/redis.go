package db

import (
	"github.com/gomodule/redigo/redis"
)

func RedisConnect() redis.Conn {
	c, err := redis.DialURL("redis://redis:6379/1")
	if err != nil {
		panic(err)
	}
	return c
}
