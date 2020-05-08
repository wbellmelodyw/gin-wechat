package cache

import (
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/utils"
	"time"
)

type weCache struct {
	cache *utils.Redis
}

func NewCache() *weCache {
	cache := utils.Get("DEFAULT") //1小时
	return &weCache{cache: cache}
}

func (we *weCache) Get(key string) interface{} {
	s := we.cache.Get(key)
	logger.Module("wechat").Sugar().Info("key", s)
	if s == "" {
		return nil
	}
	return s
}
func (we *weCache) Set(key string, val interface{}, timeout time.Duration) error {
	_ = we.cache.Set(key, val, timeout)
	return nil
}

func (we *weCache) IsExist(key string) bool {
	result := we.cache.Exists(key)
	return result == 1
}
func (we *weCache) Delete(key string) error {
	_ = we.cache.Del(key)
	return nil
}
