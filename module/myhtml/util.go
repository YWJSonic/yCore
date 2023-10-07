package myhtml

import (
	"bytes"
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

// 取得節點附加參數
//
// @params attrKey 附加參數關鍵字
//
// @params []html.Attribute 附加參數列表
//
// @return string 附加參數資料, PS: 預設為空字串
func GetAttr(attrKey string, attrs []html.Attribute) string {
	for _, attr := range attrs {
		if attr.Key == attrKey {
			return attr.Val
		}
	}
	return ""
}

// 取得所有子物件內文
//
// @params *TokenObjSub 起始解點
//
// @return string 內文
func GetContext(node *TokenObjSub) string {
	res := bytes.Buffer{}
	if node.Res.Type == html.TextToken {
		res.WriteString(node.Res.Data)
		return res.String()
	}

	for _, subnode := range node.SubRes {
		if subnode.Res.Type == html.TextToken {
			res.WriteString(subnode.Res.Data)
		} else {
			subContext := GetContext(subnode)
			res.WriteString(subContext)
		}
	}

	return res.String()
}
