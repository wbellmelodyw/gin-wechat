package db

import (
	"fmt"
	"github/wbellmelodyw/gin-wechat/model"
	"github/wbellmelodyw/gin-wechat/translate"
	"golang.org/x/text/language"
	"strings"
	"testing"
)

func TestConn(t *testing.T) {
	google := translate.GetGoogle(language.Chinese, language.English)
	text, err := google.Text("你好")
	if err != nil {
		t.Error(err)
	}
	word := new(model.Word)
	word.SrcContent = "你好"
	word.DstContent = text.Mean
	attrs := make([]string, 2)
	for name, attr := range text.Attr {
		attrs = append(attrs, name+strings.Join(attr, ";"))
	}
	word.DstAttr = strings.Join(attrs, "\n")
	explain := make([]string, 2)
	for name, e := range text.Attr {
		explain = append(explain, name+strings.Join(e, ";"))
	}
	word.DstExplain = strings.Join(explain, "\n")
	word.DstExample = strings.Join(text.Example, "\n")
	row, err := weChatDB.Insert(word)
	if row == 0 || err != nil {
		fmt.Println(row)
		fmt.Println(err)
	}
	fmt.Println(row)
	fmt.Println(err)
}
