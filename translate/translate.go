package translate

import (
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"github/wbellmelodyw/gin-wechat/logger"
	"strconv"
	"unicode/utf8"

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
	Mean    string              `json:"mean"`    //词意
	Attr    map[string][]string `json:"attr"`    //词性
	Explain map[string][]string `json:"explain"` //解释
	Example []string            `json:"example"` //造句
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
	texts.Attr = make(map[string][]string, 2)
	result := gjson.Get(rspJson, "1")
	for _, attrs := range result.Array() {
		if attrs.Get("0").String() != "" {
			for _, attr := range attrs.Get("1").Array() {
				texts.Attr[attrs.Get("0").String()] = append(texts.Attr[attrs.Get("0").String()], attr.String())
			}
		}
	}
	//解释
	result = gjson.Get(rspJson, "12")
	//wordExplain := "解释:" //少的用+就行 多才用 strings.builder
	texts.Explain = make(map[string][]string, int(result.Num))
	//logger.Module("test").Sugar().Info("Num", result.Num)
	for _, attrs := range result.Array() {
		attrName := attrs.Get("0").String()
		if attrName != "" {
			for _, attr := range attrs.Get("1").Array() {
				texts.Explain[attrName] = append(texts.Explain[attrName], attr.Get("0").String())
			}
		}
	}
	//logger.Module("test").Sugar().Info("word3", texts)
	//造句
	//wordExample := "造句:"
	exampleResult := gjson.Get(rspJson, "13.0")
	texts.Example = make([]string, 0)
	for _, example := range exampleResult.Array() {
		texts.Example = append(texts.Example, example.Get("0").String())
	}
	return texts, nil
}

func (g *GoogleTranslator) Audio(text string) ([]byte, error) {
	token := GetToken(text)

	url := "https://translate.google.cn/translate_tts"
	data := map[string]string{
		"ie":      "UTF-8",
		"q":       text,
		"tl":      g.to.String(),
		"total":   "1",
		"idx":     "0",
		"textlen": strconv.Itoa(utf8.RuneCountInString(text)),
		// "dt":     []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"},
		//"oe":   "UTF-8",
		//"otf":  "1",
		//"ssel": "0",
		//"tsel": "0",
		//"kc":   "7",
		"tk":     token,
		"client": "webapp", //"gtx",
		"prev":   "input",
	}
	client := resty.New()
	res, err := client.R().SetQueryParams(data).Get(url)
	if err != nil {
		logger.Module("audio").Sugar().Error("read", err)
		return []byte{}, err
	}
	logger.Module("audio").Sugar().Info("read", res.StatusCode())
	buffer := make([]byte, 40960)
	_, err = res.RawBody().Read(buffer)
	if err != nil {
		logger.Module("audio").Sugar().Error("read", err)
	}
	return buffer, nil
}
