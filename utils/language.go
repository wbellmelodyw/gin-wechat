package utils

import (
	"golang.org/x/text/language"
	"unicode"
)

func GetLanguageTag(s string) (from, to language.Tag) {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return language.Chinese, language.English
		}
	}
	return language.English, language.Chinese
}
