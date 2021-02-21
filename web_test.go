package main_test

import (
	"fmt"
	"testing"
	"yangServer/format/web"
	"yangServer/net"
)

func Test_Web(t *testing.T) {
	url := "https://www.alphapolis.co.jp/manga/official"
	doc, err := net.GetWebPackage(url)
	if err != nil {
		fmt.Println(err)
		return
	}
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
