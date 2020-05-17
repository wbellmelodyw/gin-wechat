package db

import (
	"fmt"
	"github/wbellmelodyw/gin-wechat/translate"
	"golang.org/x/text/language"
	"testing"
)

func TestConn(t *testing.T) {
	google := translate.GetGoogle(language.Chinese, language.English)
	text, err := google.Text("你好")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(text)
}
