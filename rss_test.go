package main_test

import (
	"fmt"
	"testing"
	"ycore/format/rss"
	"ycore/module/myhtml"
)

func TestRss(t *testing.T) {
	_, doc, err := myhtml.GetWebPackage("https://gnn.gamer.com.tw/rss.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	rssInfo, _ := rss.Format([]byte(doc))
	print(0, rssInfo.Node)
}

func print(count int, nodes []*rss.NodeInfo) []*rss.NodeInfo {
	if nodes == nil {
		return nil
	}

	tab := ""
	for i := 0; i < count; i++ {
		tab += "\t"
	}
	for _, node := range nodes {
		if node.Name == "title" {
			fmt.Printf("%s<%s>%s\n", tab, node.Name, node.Value)
		}
		print(count+1, node.IncludeNode)
		// fmt.Printf("</%s>\n", node.Name)
	}

	return nodes
}
