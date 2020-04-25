package translate

import (
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"github/wbellmelodyw/gin-wechat/logger"

	//"github/wbellmelodyw/gin-wechat/logger"
	"golang.org/x/text/language"
	"net/url"
	//"strings"
)

type GoogleTranslator struct {
	form language.Tag //来源
	to   language.Tag //要翻译成
}

type Text struct {
	Mean    string              //词意
	Attr    map[string][]string //词性
	Explain map[string][]string //解释
	Example map[string][]string //造句
}

func GetGoogle(form, to language.Tag) *GoogleTranslator {
	return &GoogleTranslator{
		form,
		to,
	}
}

func (g *GoogleTranslator) Text(text string) (*Text, error) {
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
	//提取翻译
	texts := new(Text)
	if err != nil {
		return texts, err
	}
	//texts := make([]string, 0)
	rspJson := r.String()
	//词意
	//result := gjson.Get(rspJson, "0.0")
	wordMean := "词意:" + gjson.Get(rspJson, "0.0.0").String()
	texts.Mean = wordMean
	//词性
	result := gjson.Get(rspJson, "1")
	//wordAtr := "词性:" //少的用+就行 多才用 strings.builder
	for _, attrs := range result.Array() {
		//texts.attr["词性"] = append(texts.attr["词性"],)
		//wordAtr += attrs.Get("0").String() + ":"
		if attrs.Get("0").String() != "" {
			//texts.attr[attrs.Get("0").String()] = make([]string, 0)
			//for _, attr := range attrs.Get("1").Array() {
			//	texts.attr[attrs.Get("0").String()] = append(texts.attr[attrs.Get("0").String()], attr.String())
			//}
			logger.Module("test").Sugar().Info("word3", attrs.Get("0").String())

		}
	}
	//texts = append(texts, wordAtr)
	//解释
	//result = gjson.Get(rspJson, "12")
	//wordExplain := "解释:" //少的用+就行 多才用 strings.builder
	//for _, attrs := range result.Array() {
	//	wordExplain += attrs.Get("0").String() + ":"
	//	for _, attr := range attrs.Get("1").Array() {
	//		wordExplain += attr.Get("0").String() + "|"
	//	}
	//	wordExplain = strings.TrimRight(wordExplain, "|")
	//	wordExplain += ";"
	//}
	//logger.Module("test").Sugar().Info("word3", wordExplain)
	//texts = append(texts, wordExplain)
	////造句
	//wordExample := "造句:"
	//for _, example := range gjson.Get(rspJson, "13.0").Array() {
	//	wordExample += example.Get("0").String() + "|"
	//}
	//texts = append(texts, wordExample)

	return texts, nil
}

//func Audio(text string, from language.Tag, to language.Tag) (string, error) {
//
//}
