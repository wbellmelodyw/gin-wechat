package utils

import (
	"github.com/go-redis/redis/v7"
	"github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/logger"
	"sync"
	"time"
)

//new redis collection
var rs = make(map[string]*Redis, 4)

//锁
var mu = sync.RWMutex{}

// Redis 客户端对象
type Redis struct {
	client *redis.Client
	name   string
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:53
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func Get(name string) *Redis {
	mu.Lock()
	r := rs[name]
	mu.Unlock()

	if r != nil {
		return r
	}

	c := redis.NewClient(&redis.Options{
		Addr:     config.MustGetString("REDIS_"+name+"_HOST", "redis:6379"),
		Password: config.MustGetString("REDIS_"+name+"_PASSWORD", ""), // no password set
		DB:       0,                                                   // use default DB
	})

	r = &Redis{
		client: c,
		name:   name,
	}
	mu.Lock()
	rs[name] = r
	mu.Unlock()

	return r
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) TTL(key string, time time.Duration, values ...interface{}) time.Duration {
	durationCmd := r.client.TTL(key)
	return durationCmd.Val()
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) Get(key string) string {
	stringCmd := r.client.Get(key)
	return stringCmd.String()
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) Set(key string, value interface{}, expired time.Duration) string {
	stringCmd := r.client.Set(key, value, expired)
	return stringCmd.Val()
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) HGet(key string, field string) string {
	stringCmd := r.client.HGet(key, field)
	return stringCmd.Val()
}

func (r *Redis) HGetAll(key string) (result map[string]string) {
	stringCmd := r.client.HGetAll(key)
	result, err := stringCmd.Result()
	if err != nil {
		logger.Module("biz-redis-info").Sugar().Error("hash get all error:", err)
		return
	}
	return
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) HDel(key string, fields ...string) int64 {
	intCmd := r.client.HDel(key, fields...)
	return intCmd.Val()
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) Del(key string, fields ...string) int64 {
	intCmd := r.client.Del(key)
	return intCmd.Val()
}

// @title
// @description
// @author        jeffrey  2020-04-08 10:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) HSetEx(key string, time time.Duration, values ...interface{}) int64 {
	intCmd := r.client.HSet(key, values...)
	r.client.Expire(key, time)
	return intCmd.Val()
}

func (r *Redis) HSet(key string, time time.Duration, values ...interface{}) int64 {
	intCmd := r.client.HSet(key, values...)
	return intCmd.Val()
}

// 下面是redis set 操作
// 下面是redis set 操作
// 下面是redis set 操作

// @title         集合操作（set add）
// @description
// @author        jeffrey  2020-04-08 15:48
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) SAdd(key string, members ...interface{}) int64 {
	intCmd := r.client.SAdd(key, members...)
	return intCmd.Val()
}

// @title         集合操作（set add）
// @description   可设置过期时间
// @author        jeffrey  2020-04-08 15:49
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) SAddEx(key string, time time.Duration, members ...interface{}) int64 {
	intCmd := r.client.SAdd(key, members...)
	r.client.Expire(key, time)
	return intCmd.Val()
}

// @title         SRemove
// @description   移除多个元素
// @author        jeffrey  2020-04-08 16:15
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) SRemove(key string, members ...interface{}) int64 {
	intCmd := r.client.SRem(key, members...)
	return intCmd.Val()
}

// @title         SMember
// @description	  获取集合下面所有元素
// @author        jeffrey  2020-04-08 16:16
// @param         输入参数名        参数类型         "解释"
// @return        返回参数名        参数类型         "解释"

func (r *Redis) SMember(key string) []string {
	stringSliceCmd := r.client.SMembers(key)
	return stringSliceCmd.Val()
}

func (r *Redis) Exists(key string) int64 {
	IntCmd := r.client.Exists(key)
	return IntCmd.Val()
}
