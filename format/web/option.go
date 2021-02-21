package web

import (
	"fmt"
	"net/url"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func FindByTree(option FindOption, tree *HtmlTree) []*html.Node {
	nodes := []*html.Node{}
	if option.Key == "" || option.Value == "" {
		return nodes
	}

	nodes = loop(tree.Root, option, 0)
	return nodes
}

func FindByNode(option FindOption, node *html.Node) []*html.Node {
	nodes := []*html.Node{}
	if option.Key == "" || option.Value == "" {
		return nodes
	}

	nodes = loop(node, option, 0)
	return nodes
}

type FindOption struct {
	Key      string
	Value    string
	MaxDepth int
	IsPrint  bool
}

func loop(node *html.Node, option FindOption, currentDepth int) []*html.Node {
	var findNode []*html.Node // 符合條件的Node

	// 深度限制
	if option.MaxDepth > 0 {
		if currentDepth >= option.MaxDepth {
			return findNode
		}
	}

	// 深度尋找
	if node.FirstChild != nil {
		currentDepth++
		nodes := loop(node.FirstChild, option, currentDepth)
		findNode = append(findNode, nodes...)
		currentDepth--
	}

	// 廣度尋找
	if node.NextSibling != nil {
		nodes := loop(node.NextSibling, option, currentDepth)
		findNode = append(findNode, nodes...)
	}

	// 條件比對
	for _, attr := range node.Attr {
		if attr.Key == option.Key && attr.Val == option.Value {
			findNode = append(findNode, node)
		}
	}

	return findNode
}

func PrintChild(currentDepth int, node *html.Node, option FindOption, webInfo WebInfo) {

	// 深度限制
	if option.MaxDepth > 0 {
		if currentDepth >= option.MaxDepth {
			return
		}
	}
	tab := ""
	if option.IsPrint {
		for i := 0; i < currentDepth; i++ {
			tab += " "
		}
		if node.Type == html.CommentNode {
			fmt.Println(node)
		} else if node.Type != html.TextNode {
			if node.DataAtom == atom.A {
				fmt.Printf("%s <%s %s>\n", tab, node.Data, PrintAttr(node.Attr, webInfo))
			} else {
				fmt.Printf("%s <%s %s>\n", tab, node.Data, PrintAttr(node.Attr, webInfo))
			}
		}
	}
	// 深度尋找
	if node.FirstChild != nil {
		currentDepth++
		PrintChild(currentDepth, node.FirstChild, option, webInfo)
		currentDepth--
	}

	if option.IsPrint {
		if node.Type != html.TextNode {
			fmt.Printf("%s </%s>\n", tab, node.Data)
		}
	}

	// 廣度尋找
	if node.NextSibling != nil {
		PrintChild(currentDepth, node.NextSibling, option, webInfo)
	}
}

type WebInfo interface {
	Url() string
}

func PrintAttr(attrs []html.Attribute, webInfo WebInfo) string {
	var attrsStr string
	for _, attr := range attrs {
		switch attr.Key {
		case "href":
			attr.Val = FormatHref(attr.Val, webInfo)
		case "src":
			attr.Val = FormatSrc(attr.Val, webInfo)

		}
		attrsStr += fmt.Sprintf("%s=\"%s\" ", attr.Key, attr.Val)
	}
	return attrsStr
}

func FormatHref(href string, webInfo WebInfo) string {
	if href[:4] != "http" {
		u, err := url.Parse(webInfo.Url())
		if err != nil {
			panic(href)
		}
		u.Path = href
		href = u.String()
	}
	return href
}
func FormatSrc(src string, webInfo WebInfo) string {

	if src[:4] != "http" {
		u, err := url.Parse(webInfo.Url())
		if err != nil {
			panic(src)
		}

		switch {
		case src[:2] == "./":
			u.Path = src[2:]
		case src[:4] != "http":
			u.Path = src
		}
		src, _ = url.PathUnescape(u.String())
	}

	return src
}
