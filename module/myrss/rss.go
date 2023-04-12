package myrss

import (
	"bytes"
	"encoding/xml"

	"github.com/YWJSonic/ycore/dao/queue"
	"github.com/YWJSonic/ycore/module/mylog"
)

type NodeType uint8

type RssInfo struct {
	Node       []*NodeInfo
	RssVersion string
	XmlVerSion string
	Encoding   string
}

type NodeInfo struct {
	Type        NodeType
	Name        string
	Tag         []xml.Attr
	IncludeNode []*NodeInfo
	Value       string
}

func Format(pageData []byte) (*RssInfo, error) {
	rss := &RssInfo{}

	decoder := xml.NewDecoder(bytes.NewBuffer(pageData))
	var priveNode *NodeInfo
	var currentNode *NodeInfo
	qu := queue.Queue{}
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {

		case xml.ProcInst:
			// 創建當前節點
			node := &NodeInfo{}
			node.Type = NodeType_ProcInst
			node.Name = se.Target
			node.Value = string(se.Inst)
			rss.Node = append(rss.Node, node)
			// 執行對應的操作
			procInstOption(rss, node)

		case xml.StartElement:
			// 前一個未結束的節點才會被放入層級堆疊
			if currentNode != nil {
				if priveNode != nil {
					qu.Push(priveNode)
				}
				priveNode = currentNode
			}
			// 創建當前節點
			node := &NodeInfo{}
			node.Type = NodeType_Element
			node.Name = se.Name.Local
			node.Tag = se.Attr

			if currentNode == nil && priveNode == nil {
				rss.Node = append(rss.Node, node)
			}

			// 暫存當前節點
			currentNode = node

			// 節點關聯
			if priveNode != nil {
				priveNode.IncludeNode = append(priveNode.IncludeNode, node)
			}
		case xml.EndElement:
			currentNode = nil
			if priveNode != nil && priveNode.Name == se.Name.Local {
				if node := qu.Pop(); node != nil {
					priveNode = node.(*NodeInfo)
				}
			}
		case xml.CharData:
			if currentNode == nil {
				continue
			}
			if currentNode != nil {
				currentNode.Value = string(se)
				// 執行對應的操作
				elementOption(rss, currentNode)
			}

		default:
			mylog.Info(se)
		}
	}

	return rss, nil
}
