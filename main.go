package main

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github/wbellmelodyw/gin-wechat/hello"
	"github/wbellmelodyw/gin-wechat/translate"
	"golang.org/x/text/language"
	"time"
)

func main() {

	text := "hello,world"
	//fmt.Println(language.Chinese.String())
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From:  "en",
			To:    "zh",
			Delay: time.Second,
		})
	if err != nil {
		panic(err)
	}
	fmt.Printf("en: %s | zh: %s \n", text, translated)

	text2 := "In the good old days of computing when  memory was expensive and processing power was at premium, hacking on bits directly was the preferred (in some cases the only) way to process information. Today, direct bit manipulation is still crucial in many computing use cases such as low-level system programming, image processing, cryptography, etc."
	translatedText, _ := gtranslate.Translate(text2, language.English, language.Chinese)

	fmt.Println("translated:", translatedText)
	fmt.Println(hello.Greet(translatedText))

	//test kk

	translatedText2, err := translate.Text(text2, language.English, language.Chinese)
	fmt.Println("translated2:", translatedText2)

}
