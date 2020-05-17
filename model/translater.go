package model

import "golang.org/x/text/language"

type GoogleTranslator struct {
	form language.Tag //来源
	to   language.Tag //要翻译成
}

type Text struct {
	Content string              `json:"content"` //词
	Mean    string              `json:"mean"`    //词意
	Attr    map[string][]string `json:"attr"`    //词性
	Explain map[string][]string `json:"explain"` //解释
	Example []string            `json:"example"` //造句
}
