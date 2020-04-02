package main

import (
	"fmt"
	"github/wbellmelodyw/gin-wechat/translate"
	"golang.org/x/text/language"
)

func main() {

	text := "hello,world"
	//fmt.Println(language.Chinese.String())

	//test kk
	translatedText, _ := translate.Text(text, language.English, language.Chinese)
	fmt.Println("translated:", translatedText)

	translatedText2, _ := translate.Text("rookie", language.English, language.Chinese)
	fmt.Println("translated2:", translatedText2)
}
