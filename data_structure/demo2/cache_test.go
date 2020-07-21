package demo2

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cacheFactory := &CacheFactory{}
	redis, err := cacheFactory.Create(0)
	if err != nil {
		t.Error(err)
	}
	redis.Set("redis key", "redis val")
	fmt.Println(redis.Get("redis key"))

	mem, err := cacheFactory.Create(1)
	if err != nil {
		t.Error(err)
	}
	mem.Set("mem key","mem val")
	fmt.Println(mem.Get("mem key"))
}
