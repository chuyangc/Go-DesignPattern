package SimpleFactory

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cf := CacheFactory{}
	rediscache, err := cf.Create(redis)
	if err != nil {
		t.Error(err)
	}
	rediscache.Set("redis", 1)
	fmt.Println(rediscache.Get("redis"))

	mcache, err := cf.Create(memcache)
	if err != nil {
		t.Error(err)
	}
	mcache.Set("memcache", 2)
	fmt.Println(mcache.Get("memcache"))
}
