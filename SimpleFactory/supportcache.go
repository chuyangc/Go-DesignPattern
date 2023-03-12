package SimpleFactory

import "errors"

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

// 实现Cache简单工厂
type cacheType int

const (
	redis cacheType = iota
	memcache
)

type CacheFactory struct {
}

func (cf *CacheFactory) Create(ct cacheType) (Cache, error) {
	if ct == redis {
		return &RedisCache{data: map[string]interface{}{}}, nil
	}
	if ct == memcache {
		return &MemCache{data: map[string]interface{}{}}, nil
	}
	return nil, errors.New("unknown cache type")
}
