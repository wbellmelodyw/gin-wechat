package translate

import (
	"encoding/json"
	"errors"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func Text(text string, from language.Tag, to language.Tag) (string, error) {
	token := GetToken(text)
	urll := "https://translate.google.com/translate_a/single"
	data := map[string]string{
		"client": "gtx",
		"sl":     from.String(),
		"tl":     to.String(),
		"hl":     to.String(),
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
