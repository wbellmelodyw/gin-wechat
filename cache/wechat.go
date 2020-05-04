package cache

import (
	"github.com/allegro/bigcache"
	"time"
)

type weCache struct {
	cache *bigcache.BigCache
}

func NewCache() *weCache {
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(60 * time.Minute)) //1小时
	return &weCache{cache: cache}
}

func (we *weCache) Get(key string) interface{} {
	s, _ := we.cache.Get(key)
	return string(s)
}
func (we *weCache) Set(key string, val interface{}, timeout time.Duration) error {
	return we.cache.Set(key, val.([]byte))
}
func (we *weCache) IsExist(key string) bool {
	return false
}
func (we *weCache) Delete(key string) error {
	return we.cache.Delete(key)
}
