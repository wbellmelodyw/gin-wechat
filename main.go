package main

import (
	"fmt"
	"github/wbellmelodyw/gin-wechat/config"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/translate"
	"golang.org/x/text/language"
)

func main() {

	text := "hello,world"
	//fmt.Println(language.Chinese.String())
	if err := config.Init(); err != nil {
		logger.Module("system-error").Sugar().Error("load config fail", err)
	}

	//test kk
	translatedText, _ := translate.Text(text, language.English, language.Chinese)
	fmt.Println("translated:", translatedText)

	translatedText2, _ := translate.Text("rookie", language.English, language.Chinese)
	fmt.Println("translated2:", translatedText2)
}
