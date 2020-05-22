package translate

import (
	"bytes"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"github/wbellmelodyw/gin-wechat/logger"
	"github/wbellmelodyw/gin-wechat/model"
	"os"
	"path/filepath"
	"strconv"
	"unicode/utf8"

	//"github/wbellmelodyw/gin-wechat/logger"
	"golang.org/x/text/language"
	"net/url"
	//"strings"
)

var mkdirAll = os.MkdirAll

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

func (g *GoogleTranslator) Text(text string) (*model.Text, error) {
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
	texts := new(model.Text)
	texts.Content = text
	if err != nil {
		return texts, err
	}
	//texts := make([]string, 0)
	rspJson := r.String()
	//词意
	//result := gjson.Get(rspJson, "0.0")
	wordMean := gjson.Get(rspJson, "0.0.0").String()
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
	ttsUrl := "https://translate.google.cn/translate_tts"
	data := map[string]string{
		"ie":      "UTF-8",
		"q":       text,
		"tl":      g.to.String(),
		"total":   "1",
		"idx":     "0",
		"textlen": strconv.Itoa(utf8.RuneCountInString(text)),
		"tk":      token,
		"client":  "gtx", //"gtx",
		"prev":    "input",
	}
	client := resty.New()
	res, err := client.R().SetDoNotParseResponse(true).SetQueryParams(data).Get(ttsUrl)
	defer res.RawBody().Close()
	if err != nil {
		logger.Module("audio").Sugar().Error("response error", err)
		return []byte{}, err
	}
	buffer := make([]byte, 4096)
	wBuffer := new(bytes.Buffer)

	for {
		n, err := res.RawBody().Read(buffer)
		if n == 0 || err != nil {
			logger.Module("audio").Sugar().Info("over", n)
			logger.Module("audio").Sugar().Error("over", err)
			break
		}
		wBuffer.Write(buffer[:])
	}

	return wBuffer.Bytes(), nil
}

func (g *GoogleTranslator) AudioSaveFile(text string) {
	token := GetToken(text)
	ttsUrl := "https://translate.google.cn/translate_tts"
	data := map[string]string{
		"ie":      "UTF-8",
		"q":       text,
		"tl":      g.to.String(),
		"total":   "1",
		"idx":     "0",
		"textlen": strconv.Itoa(utf8.RuneCountInString(text)),
		"tk":      token,
		"client":  "gtx", //"gtx",
		"prev":    "input",
	}
	client := resty.New()
	res, err := client.R().SetDoNotParseResponse(true).SetQueryParams(data).Get(ttsUrl)
	defer res.RawBody().Close()
	if err != nil {
		logger.Module("audio").Sugar().Error("response error", err)
	}
	buffer := make([]byte, 4096)
	createDirectoryIfMedia("/app/media/")
	file, err := os.Create("media/" + text + ".mp3")
	if err != nil {
		logger.Module("audio").Sugar().Panic("create file error", err)
	}
	for {
		n, err := res.RawBody().Read(buffer)
		if n == 0 || err != nil {
			logger.Module("audio").Sugar().Info("over", n)
			logger.Module("audio").Sugar().Error("over", err)
			break
		}
		file.Write(buffer[:])
	}
}

func createDirectoryIfMedia(connection string) {
	logger.Module("create").Sugar().Info("file", filepath.Dir(connection))
	if _, err := os.Stat(filepath.Dir(connection)); os.IsNotExist(err) {
		if err := mkdirAll(filepath.Dir(connection), 0777); err != nil {
			panic(err)
		}
	}
}
