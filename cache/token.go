package cache

import (
	"github.com/allegro/bigcache"
	"time"
)

var cache *bigcache.BigCache

func init() {
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(60 * time.Minute)) //1小时
}

func GetToken(key string) (tkk string, ok bool) {
	value, err := cache.Get(key)
	if err != nil {
		return
	}
	return string(value), true
}

func SetToken(key, tkk string) error {
	return cache.Set(key, []byte(tkk))
}
