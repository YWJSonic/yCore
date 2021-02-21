package rss

import (
	"encoding/xml"
	"strings"
)

func replaceHtmlStr(input string) string {
	return replacer.Replace(input)
}

func elementOption(rss *RssInfo, node *NodeInfo) {
	if node.Type != NodeType_Element {
		return
	}

	switch node.Name {
	case "rss":
		for _, attr := range node.Tag {
			if attr.Name.Local == "version" {
				rss.RssVersion = attr.Value
			}
		}
	case "description":
		str := replaceHtmlStr(node.Value)
		node.Value = str
	}
}

func procInstOption(rss *RssInfo, node *NodeInfo) {
	if node.Type != NodeType_ProcInst {
		return
	}

	switch node.Name {
	case "xml":
		procInstFormat_xml(node)
		for _, tags := range node.Tag {
			if tags.Name.Local == "version" {
				rss.XmlVerSion = tags.Value
			} else if tags.Name.Local == "encoding" {
				rss.Encoding = tags.Value
			}
		}
	}
}

func procInstFormat_xml(node *NodeInfo) {
	value := strings.ReplaceAll(node.Value, "\"", "")
	tagsStr := strings.Split(value, " ")
	for _, tagStr := range tagsStr {
		valStr := strings.Split(tagStr, "=")

		if len(valStr) < 2 {
			continue
		}

		node.Tag = append(node.Tag, xml.Attr{
			Name:  xml.Name{Local: valStr[0]},
			Value: valStr[1]})
	}
}
