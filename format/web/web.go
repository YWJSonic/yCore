package web

import (
	"strings"

	"golang.org/x/net/html"
)

func Format(pageData []byte, url string) *HtmlTree {
	tree := &HtmlTree{
		url: url,
	}
	node, _ := html.Parse(strings.NewReader(string(pageData)))
	tree.Root = node
	return tree
}

func DetailFormat() {
	// z := html.NewTokenizer(strings.NewReader(string(pageData)))
	// depth := 0
	// for {
	// 	tt := z.Next()
	// 	err := z.Err()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	switch tt {
	// 	case html.DoctypeToken:
	// 		token := z.Token()
	// 		fmt.Println(token)
	// 	case html.TextToken:
	// 		token := z.Token()
	// 		fmt.Println(token)
	// 		// fmt.Println(z.TagName())
	// 		// fmt.Println(string(z.Text()))
	// 		if depth > 0 {
	// 			// emitBytes should copy the []byte it receives,
	// 			// if it doesn't process it immediately.
	// 			// emitBytes(z.Text())
	// 		}

	// 	case html.StartTagToken:
	// 		token := z.Token()
	// 		fmt.Println(token)
	// 		tn, _ := z.TagName()
	// 		if len(tn) == 1 && tn[0] == 'a' {
	// 			depth++
	// 		}
	// 	case html.EndTagToken:
	// 		token := z.Token()
	// 		fmt.Println(token)
	// 		depth--
	// 	}
	// }
}
