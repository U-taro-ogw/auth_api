package handlers

import (
	"github.com/gomodule/redigo/redis"
)

type RedisHandler struct {
  Redis redis.Conn
}

func (h *RedisHandler) Get(key string) string {
	res, _ := redis.String(h.Redis.Do("GET", key))
	return res
}

func (h *RedisHandler) Set(key, value string) {
	h.Redis.Do("SET", key, value)
}