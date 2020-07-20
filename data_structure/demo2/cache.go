package demo2

import "errors"

type Cache interface {
	Set(key, value string)
	Get(key string) string
}

type RedisCache struct {
	data map[string]string
}

func (redis *RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return redis.data[key]
}

type MemCache struct {
	data map[string]string
}

func (mem *MemCache) Set(key, val string) {
	mem.data[key] = val
}

func (mem *MemCache) Get(key string) string {
	return mem.data[key]
}

type CacheType int

const (
	redis CacheType = iota
	mem
)

type CacheFactory struct {
}

func (factory *CacheFactory) Create(cacheType CacheType) (Cache, error) {
	if cacheType == redis {
		return &RedisCache{data: map[string]string{}}, nil
	}

	if cacheType == mem {
		return &MemCache{data: map[string]string{}},nil
	}
	return nil,errors.New("error cache type")
}
