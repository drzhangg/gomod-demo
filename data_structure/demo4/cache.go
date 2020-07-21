package demo4

type Cache interface {
	Set(key string, val interface{})
	Get(key string) interface{}
}

type Redis struct {
	data map[string]interface{}
}

func (r *Redis) Set(key string, val interface{}) {
	r.data[key] = val
}

func (r *Redis) Get(key string) interface{} {
	return r.data[key]
}

type MemCache struct {
	data map[string]interface{}
}

func (m *MemCache) Set(key string, val interface{}) {
	m.data[key] = val
}
func (m *MemCache) Get(key string) interface{} {
	return m.data[key]
}

type CacheFactory interface {
	Create() Cache
}

type RedisFactory struct {
}

func (r *RedisFactory) Create() Cache {
	return &Redis{}
}

type MemFactory struct {
}

func (m *MemFactory) Create() Cache {
	return &MemCache{}
}
