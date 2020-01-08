package db

import (
	"github.com/gomodule/redigo/redis"
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	return c
}

func Get(key string, c redis.Conn) string {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res
}