package myhtml

// Filter Operation
const (
	FilterOperation_GetContent     int = 1 // 取得目標標籤內文
	FilterOperation_GetSubToken    int = 2 // 取得目標標籤內的子標籤
	FilterOperation_GetSubcContent int = 3 // 取得目標標籤內的子內文

)

// var typeMap map[html.NodeType]string = map[html.NodeType]string{
// 	html.ErrorNode:    "ErrorNode",
// 	html.TextNode:     "TextNode",
// 	html.DocumentNode: "DocumentNode",
// 	html.ElementNode:  "ElementNode",
// 	html.CommentNode:  "CommentNode",
// 	html.DoctypeNode:  "DoctypeNode",
// 	html.RawNode:      "RawNode",
// }
