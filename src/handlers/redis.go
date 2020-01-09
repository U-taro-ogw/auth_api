package handlers

import (
	"github.com/gomodule/redigo/redis"
)

type RedisHandler struct {
  Redis redis.Conn
}

func (h *RedisHandler) Get(key string) string {
	res, err := redis.String(h.Redis.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res
}