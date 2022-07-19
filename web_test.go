package main_test

import (
	"fmt"
	"testing"
	"ycore/format/web"
	"ycore/module/myhtml"
)

func Test_Web(t *testing.T) {
	url := "https://union.591.com.tw/stats/event?c=page-pc&a=page-4&l=page-1&_u=q13broe0dcg04oh5fjrh25gom4"
	header, doc, err := myhtml.GetWebPackage(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range header {
		fmt.Println(k, v)
	}
	fmt.Println(string(doc))
	htmlTree := web.Format(doc, url)

	nods := web.FindByNode(web.FindOption{
		Key:      "class",
		Value:    "section recently-mangas",
		MaxDepth: 0,
	}, htmlTree.Root)

	nods = web.FindByNode(web.FindOption{
		Key:      "class",
		Value:    "wrap-mangas-list",
		MaxDepth: 0,
	}, nods[0])

	nods = web.FindByNode(web.FindOption{
		Key:      "class",
		Value:    "mangas-list",
		MaxDepth: 0,
	}, nods[0])

	web.PrintChild(
		0,
		nods[0],
		web.FindOption{
			MaxDepth: 0,
			IsPrint:  true,
		},
		htmlTree)
	// for _, node := range nods {
	// 	fmt.Print(node.Data)
	// 	for _, attr := range node.Attr {
	// 		fmt.Printf(" Key:%s Value:%s ", attr.Key, attr.Val)
	// 	}
	// 	fmt.Print(node.FirstChild.Data)
	// 	fmt.Println("")
	// }
}
