package translate

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"github/wbellmelodyw/gin-wechat/logger"
	"golang.org/x/text/language"
	"net/url"
)

type GoogleTranslator struct {
	form language.Tag //来源
	to   language.Tag //要翻译成
}

func GetGoogle(form, to language.Tag) *GoogleTranslator {
	return &GoogleTranslator{
		form,
		to,
	}
}

func (g *GoogleTranslator) Text(text string) (string, error) {
	token := GetToken(text)

	urll := "https://translate.google.com/translate_a/single"
	data := map[string]string{
		"client": "gtx",
		"sl":     g.form.String(),
		"tl":     g.to.String(),
		"hl":     g.to.String(),
		// "dt":     []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"},
		"ie":   "UTF-8",
		"oe":   "UTF-8",
		"otf":  "1",
		"ssel": "0",
		"tsel": "0",
		"kc":   "7",
		"q":    text,
		"tk":   token,
	}
	client := resty.New()

	r, err := client.R().SetQueryParams(data).
		SetQueryParamsFromValues(url.Values{
			"dt": []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"},
		}).Get(urll)

	var resp []interface{}
	//提取翻译
	//texts := make([]string,5)
	result := gjson.Get(r.String(), "0.0")
	var a string
	for _, name := range result.Array() {
		a += name.String()
	}
	logger.Module("test").Sugar().Info("resp1", a)

	err = json.Unmarshal(r.Body(), &resp)
	if err != nil {
		return "", err
	}
	responseText := ""
	logger.Module("test").Sugar().Info("resp", resp)

	//翻译
	for _, obj := range resp[0].([]interface{}) {
		if len(obj.([]interface{})) == 0 {
			break
		}

		t, ok := obj.([]interface{})[0].(string)
		if ok {
			responseText += t
		}
	}
	//词性

	return responseText, nil
}

//func Audio(text string, from language.Tag, to language.Tag) (string, error) {
//
//}
