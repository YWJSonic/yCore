package myhtml

import (
	"io"
	"log"

	"github.com/YWJSonic/ycore/types"
	"golang.org/x/net/html"
)

// 取得符合的底下所有仔物件
type FilterObjnewSub struct {
	FiltAttrs []html.Attribute // 標籤篩選條件
	Res       []*TokenObjSub   // 標籤結構
	Operation []int            // 控制選項 暫時用數字處理
}

type TokenObjSub struct {
	NodeDepth int
	NodeId    int
	Previous  *TokenObjSub   // 前一個
	Res       *html.Token    // 當前
	SubRes    []*TokenObjSub // 子標籤結構
}

// 建立全物件索引, 並索引符合條件的物件
//
// @params *html.Tokenizer 頁面資料
//
// @params map[types.TokenTypeName][]*FilterObjnewSub 塞選器物件
//
// @params func(nodeDepth int, next, current, previous **TokenObjSub) (int, bool) 頁面修正方法
func HtmlLoopFilterLevelSub(tokenizer *html.Tokenizer, filterMap map[types.TokenTypeName][]*FilterObjnewSub, pageFix func(nodeDepth int, next, current, previous **TokenObjSub) (int, bool)) {

	var previous, current *TokenObjSub
	nodeDepth := 0 // 層級深度 0~n
	nodeId := 0    // 節點編號 1~n
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		token := tokenizer.Token()

		// 調整標籤類型 *有可能開發者未填入結束標籤
		htmlType := token.Type
		if token.Type == html.StartTagToken {
			switch token.Data {
			case "meta", "link", "input", "img":
				htmlType = html.SelfClosingTagToken
			}
		}

		switch htmlType {
		case html.StartTagToken:
			nodeId++

			next := &TokenObjSub{
				NodeDepth: nodeDepth,
				NodeId:    nodeId,
				Res:       &token,
			}

			// if previous != nil {
			// 	fmt.Println(previous.Res.Data, previous.Res.Attr)
			// }
			// if current != nil {
			// 	fmt.Println(current.Res.Data, current.Res.Attr)
			// }
			// fmt.Println(next.Res.Data, next.Res.Attr)
			// fmt.Println("----------------------------------")

			// 修正節點錯誤
			if pageFix != nil {
				if newNodeDepth, isFix := pageFix(nodeDepth, &next, &current, &previous); isFix {
					nodeDepth = newNodeDepth

					// ---------------------
					// if previous != nil {
					// 	fmt.Println(previous.Res.Data, previous.Res.Attr)
					// }
					// if current != nil {
					// 	fmt.Println(current.Res.Data, current.Res.Attr)
					// }
					// fmt.Println(next.Res.Data, next.Res.Attr)
					// fmt.Println("--------------Fix--------------------")
				}
			}

			if current == nil {
				current = next
			} else {
				next.Previous, previous = current, current
				current = next
			}

			if previous != nil {
				previous.SubRes = append(previous.SubRes, current)
			}

			// 辨識篩選目標
			targets, ok := filterMap[current.Res.Data]
			if ok {
				for _, target := range targets {
					if AttrsCompare(target.FiltAttrs, current.Res) {
						target.Res = append(target.Res, current)
					}
				}
			}
			nodeDepth++

		case html.EndTagToken:

			if previous != nil {
				current, previous = current.Previous, current.Previous.Previous
			}
			nodeDepth--

		case html.TextToken:
			// 無內文結束
			if IsTextEnd(token.Data) {
				break
			}

			next := &TokenObjSub{
				NodeDepth: nodeDepth,
				NodeId:    nodeId,
				Res:       &token,
			}
			current.SubRes = append(current.SubRes, next)

		case html.SelfClosingTagToken: // 自行封閉節點例如 <meta />
			nodeId++
			next := &TokenObjSub{
				NodeDepth: nodeDepth,
				NodeId:    nodeId,
				Res:       &token,
			}
			if current != nil {
				current.SubRes = append(current.SubRes, next)
			}

			// 辨識篩選目標
			targets, ok := filterMap[next.Res.Data]
			if !ok {
				break
			}

			for _, target := range targets {
				if AttrsCompare(target.FiltAttrs, next.Res) {
					target.Res = append(target.Res, next)
				}
			}
		}
	}
}
