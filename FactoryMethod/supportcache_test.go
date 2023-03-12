package FactoryMethod

import (
	"fmt"
	"testing"
)

func TestRedisFactory_Create(t *testing.T) {
	rf := RedisFactory{}
	rc := rf.Create()
	rc.Set("redis", 1)
	fmt.Println(rc.Get("redis"))
}

func TestMemcacheFactory_Create(t *testing.T) {
	mf := MemcacheFactory{}
	mc := mf.Create()
	mc.Set("memcache", 2)
	fmt.Println(mc.Get("memcache"))
}
