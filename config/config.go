package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() (err error) {
	path := "/Users/ace/code/vgo/gin-wechat"
	viper.SetConfigName("gin")
	fmt.Println(path)
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	return
}

// GetString 获取字符串配置
func GetString(key string) string {
	return viper.GetString(key)
}
