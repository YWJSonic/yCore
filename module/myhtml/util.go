package myhtml

import (
	"strings"

	"golang.org/x/net/html"
)

// 目標 attr 是否存在 token 內
func AttrCompare(targetAttr html.Attribute, token *html.Token) bool {
	for _, attr := range token.Attr {
		if targetAttr == attr {
			return true
		}
	}
	return false
}

// 目標 attrs 是否全部都在目標內
func AttrsCompare(targetAttrs []html.Attribute, token *html.Token) bool {
	for _, targetAttr := range targetAttrs {
		if !AttrCompare(targetAttr, token) {
			return false
		}
	}
	return true
}

func IsTextEnd(text string) bool {
	// 結束符號
	if text == "\n" {
		return true
	} else { // 結束符號含其他格式符號
		contant := strings.ReplaceAll(text, "\t", "")
		contant = strings.ReplaceAll(contant, "\r", "")
		contant = strings.ReplaceAll(contant, " ", "")
		contant = strings.ReplaceAll(contant, "\n", "")
		if contant == "" {
			return true
		}
	}
	return false
}
