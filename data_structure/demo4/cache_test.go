package demo4

import (
	"fmt"
	"testing"
)

func TestCacheFactory(t *testing.T) {
	redisFactory := RedisFactory{}
	redis := redisFactory.Create()
	redis.Set("redis", 1)
	val := redis.Get("redis")
	fmt.Println(val)
}
