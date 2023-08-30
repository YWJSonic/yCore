package myhtml

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func Test_FixPage(t *testing.T) {

	loginHtml := html.NewTokenizer(bytes.NewBuffer([]byte("")))
	filtersMap := map[string][]*FilterObjnewSub{
		"li": {
			{
				FiltAttrs: []html.Attribute{ // 收藏
					{
						Key: "class",
						Val: "collect",
					},
				},
			},
			{
				FiltAttrs: []html.Attribute{ // [押金], [最端租期], [產權登記]
					{
						Key: "class",
						Val: "clearfix",
					},
				},
			},
		},
	}

	// 修正未填寫 <li> 結束標籤問題
	pageFix := func(nodeDepth int, next, current, previous **TokenObjSub) (int, bool) {
		if (*next).Res.Data == "li" && (*current).Res.Data == "li" {
			if previous != nil {
				(*current), (*previous) = (*current).Previous, (*current).Previous.Previous
			}
			nodeDepth--
			return nodeDepth, true
		} else {
			return nodeDepth, false
		}
	}
	HtmlLoopFilterLevelSub(loginHtml, filtersMap, pageFix)
}
