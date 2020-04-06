package main

import (
	"github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/router"
	"net/http"
)

func main() {
	//text := "hello,world"
	//fmt.Println(language.Chinese.String())
	if err := config.Init(); err != nil {
		logger.Module("system-error").Sugar().Error("load config fail", err)
		panic(err)
	}
	//初始化路由
	engine := router.Create()
	//test kk
	//translatedText, _ := translate.Text(text, language.English, language.Chinese)
	//fmt.Println("translated:", translatedText)
	//
	//translatedText2, _ := translate.Text("rookie", language.English, language.Chinese)
	//fmt.Println("translated2:", translatedText2)
	addr := config.MustGetString("HOST", "localhost:8080")
	logger.Module("engine").Sugar().Panic("listen crash", http.ListenAndServe(addr, engine))
}
