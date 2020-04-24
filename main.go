package main

import (
	"fmt"
	"github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/router"
	"github/wbellmelodyw/gin-wechat/translate"
	"golang.org/x/text/language"
	"net/http"
)

func main() {
	text := "hello,world"
	traslator := translate.GetGoogle(language.English, language.Chinese)

	translatedText, _ := traslator.Text(text)
	fmt.Println("translated:", translatedText)
	//fmt.Println(language.Chinese.String())
	if err := config.Init(); err != nil {
		logger.Module("system-error").Sugar().Error("load config fail", err)
		panic(err)
	}
	//初始化路由
	engine := router.Create()
	//test kk

	//
	//translatedText2, _ := translate.Text("rookie", language.English, language.Chinese)
	//fmt.Println("translated2:", translatedText2)
	addr := config.MustGetString("HOST", "localhost:8080")
	logger.Module("engine").Sugar().Panic("listen crash", http.ListenAndServe(addr, engine))
}
