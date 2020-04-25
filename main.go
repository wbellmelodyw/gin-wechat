package main

import (
	"github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/router"
	"net/http"
)

func main() {
	if err := config.Init(); err != nil {
		logger.Module("system-error").Sugar().Error("load config fail", err)
		panic(err)
	}
	//text := "hello,world"
	//translator := translate.GetGoogle(language.English, language.Chinese)
	//translatedText, _ := translator.Text(text)
	//fmt.Println("translated:", translatedText)
	//fmt.Println(language.Chinese.String())

	//初始化路由
	engine := router.Create()
	//test kk

	//
	//translatedText2, _ := translate.Text("rookie", language.English, language.Chinese)
	//fmt.Println("translated2:", translatedText2)
	addr := config.MustGetString("HOST", "localhost:8080")
	logger.Module("engine").Sugar().Panic("listen crash", http.ListenAndServe(addr, engine))
}
