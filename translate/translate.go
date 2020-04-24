package translate

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github/wbellmelodyw/gin-wechat/logger"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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

func (g *GoogleTranslator) Text2(text string) (string, error) {
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
	}
	u, err := url.Parse(urll)
	if err != nil {
		return "", nil
	}

	parameters := url.Values{}

	for k, v := range data {
		parameters.Add(k, v)
	}
	for _, v := range []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"} {
		parameters.Add("dt", v)
	}

	parameters.Add("tk", token)
	u.RawQuery = parameters.Encode()

	var r *http.Response
	var tries = 2
	delay := time.Second

	for tries > 0 {
		r, err = http.Get(u.String())
		if err != nil {
			if err == http.ErrHandlerTimeout {
				return "", errors.New("bad network, please check your internet connection")
			}
			return "", err
		}

		if r.StatusCode == http.StatusOK {
			break
		}

		if r.StatusCode == http.StatusForbidden {
			tries--
			time.Sleep(delay)
		}
	}

	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var resp []interface{}

	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {
		return "", err
	}

	responseText := ""
	logger.Module("test").Sugar().Info("resp", resp)
	for _, obj := range resp[0].([]interface{}) {
		if len(obj.([]interface{})) == 0 {
			break
		}

		t, ok := obj.([]interface{})[0].(string)
		if ok {
			responseText += t
		}
	}

	return responseText, nil
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
	}
	client := resty.New()
	//u, err := url.Parse(urll)
	//if err != nil {
	//	return "", nil
	//}
	r, err := client.R().SetQueryParams(data).SetQueryParam("tk", token).
		SetQueryParamsFromValues(url.Values{
			"dt": []string{"at", "bd", "ex", "ld", "md", "qca", "rw", "rm", "ss", "t"},
		}).Get(urll)

	//parameters.Add("tk", token)
	//u.RawQuery = parameters.Encode()

	//var r *http.Response
	//var tries = 2
	//delay := time.Second
	//
	//for tries > 0 {
	//	r, err = http.Get(u.String())
	//	if err != nil {
	//		if err == http.ErrHandlerTimeout {
	//			return "", errors.New("bad network, please check your internet connection")
	//		}
	//		return "", err
	//	}
	//
	//	if r.StatusCode == http.StatusOK {
	//		break
	//	}
	//
	//	if r.StatusCode == http.StatusForbidden {
	//		tries--
	//		time.Sleep(delay)
	//	}
	//}

	//raw, err := ioutil.ReadAll(r.Body())
	//if err != nil {
	//	return "", err
	//}

	var resp []interface{}

	err = json.Unmarshal(r.Body(), &resp)
	if err != nil {
		return "", err
	}

	responseText := ""
	logger.Module("test").Sugar().Info("resp", resp)
	for _, obj := range resp[0].([]interface{}) {
		if len(obj.([]interface{})) == 0 {
			break
		}

		t, ok := obj.([]interface{})[0].(string)
		if ok {
			responseText += t
		}
	}

	return responseText, nil
}

//func Audio(text string, from language.Tag, to language.Tag) (string, error) {
//
//}
