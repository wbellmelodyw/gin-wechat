package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getTestRedis() *Redis {
	return Get("TEST")
}

func TestRedis_HSetEx(t *testing.T) {
	r := getTestRedis()
	row := r.HSetEx("a", 2*time.Second, "u_1", "read")
	assert.Equal(t, int64(1), row, "插入成功")
	value := r.HGet("a", "u_1")
	assert.Equal(t, "read", value, "插入成功")
	time.Sleep(2 * time.Second)
	v2 := r.HGet("a", "u_1")
	assert.Equal(t, "", v2, "key没过期")
}

//func TestRedis_SMember(t *testing.T) {
//	r := getTestRedis()
//	members := r.SMember("a")
//	fmt.Println(members)
//}
