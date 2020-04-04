package logger

import (
	"github/wbellmelodyw/gin-wechat/config"
	"testing"
)

func TestSugar(t *testing.T) {
	config.Init()
	//gin :=&gin.Context{}
	Module("test").Sugar().Error("test test test", []string{"a", "b", "c"})
}
