package web

import "golang.org/x/net/html"

var typeMap map[html.NodeType]string = map[html.NodeType]string{
	html.ErrorNode:    "ErrorNode",
	html.TextNode:     "TextNode",
	html.DocumentNode: "DocumentNode",
	html.ElementNode:  "ElementNode",
	html.CommentNode:  "CommentNode",
	html.DoctypeNode:  "DoctypeNode",
	html.RawNode:      "RawNode",
}
