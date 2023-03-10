package FactoryMethod

// Cache 定义一个接口
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

// RedisCache 定义RedisCache并实现Cache的方法
type RedisCache struct {
	data map[string]interface{}
}

func (rc *RedisCache) Set(key string, value interface{}) {
	rc.data[key] = value
}

func (rc *RedisCache) Get(key string) interface{} {
	return rc.data[key]
}

// MemCache 定义Memcache并实现Cache的方法
type MemCache struct {
	data map[string]interface{}
}

func (mc *MemCache) Set(key string, value interface{}) {
	mc.data[key] = value
}

func (mc *MemCache) Get(key string) interface{} {
	return mc.data[key]
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		data: make(map[string]interface{}),
	}
}

func NewMemcache() *MemCache {
	return &MemCache{
		data: make(map[string]interface{}),
	}
}

// CacheFactory 定义一个抽象工厂
type CacheFactory interface {
	Create() Cache
}

// RedisFactory 实现Redis工厂
type RedisFactory struct {
}

func (rf *RedisFactory) Create() Cache {
	return NewRedisCache()
}

// MemcacheFactory 定义一个抽象工厂
type MemcacheFactory struct {
}

func (mf *MemcacheFactory) Create() Cache {
	return NewMemcache()
}
