package myhtml

import "golang.org/x/net/html"

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
