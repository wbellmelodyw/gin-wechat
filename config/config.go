package config

import (
	"github.com/spf13/viper"
	"os"
)

func Init() (err error) {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	viper.SetConfigName("gin")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()
	return
}

// GetString 获取字符串配置
func MustGetString(key, defaultValue string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}

// GetString 获取字符串配置
func GetString(key string) string {
	return viper.GetString(key)
}

// GetString 获取bool配置
func GetBool(key string) bool {
	return viper.GetBool(key)
}
