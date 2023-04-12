package myhtml

import (
	"io"
	"log"

	"github.com/YWJSonic/ycore/types"
	"github.com/YWJSonic/ycore/util"

	"golang.org/x/net/html"
)

func HtmlLoop(tokenizer *html.Tokenizer, filter map[types.TokenTypeName][]*FilterObj) {

	isTarget := false
	for {
		//get the next token type
		tokenType := tokenizer.Next()

		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		if tokenType == html.StartTagToken {
			//get the token
			token := tokenizer.Token()
			if filters, ok := filter[token.Data]; ok {

				isTarget = true // 預設成功
				for _, filter := range filters {

					// 比對 token 資料是否相符
					for _, filtAttr := range filter.FiltAttrs {
						// 篩選的資料只要有一筆不存在就算失敗
						if !AttrCompare(filtAttr, token) {
							isTarget = false
							break
						}
					}

					if !isTarget {
						continue
					}

					// 找到目標 token
					filter.Res = append(filter.Res, token)
				}
			}

		}
	}
}

func HtmlLoopFilterOne(tokenizer *html.Tokenizer, filterMap map[types.TokenTypeName][]*FilterObj) {

	// 當前符合的篩選器
	targetFilters := map[*FilterObj]struct{}{}
	// 之前符合的篩選器
	lasttargetFilters := map[*FilterObj]struct{}{}

	for {
		//get the next token type
		tokenType := tokenizer.Next()

		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		var token *html.Token
		tmp := tokenizer.Token()
		token = &tmp

		// fmt.Println("token:", token.Type.String(), token.Data)

		switch tokenType {
		case html.StartTagToken:

			// 處理之前的篩選器
			for filter := range lasttargetFilters {
				isSubToken := (util.FastSearchWithInt(FilterOperation_GetSubToken, filter.Operation)) != -1

				if isSubToken {
					filter.SubRes = append(filter.SubRes, *token)
				}
			}

			// 檢查全部篩選目標
			filters, ok := filterMap[token.Data]
			if !ok {
				continue
			}

			// 相同 token 的篩選
			for _, filter := range filters {

				isTarget := true // 預設成功

				// 比對 token 資料是否相符
				for _, filtAttr := range filter.FiltAttrs {
					// 篩選的資料只要有一筆不存在就算失敗
					if !AttrCompare(filtAttr, *token) {
						isTarget = false
						break
					}
				}

				if !isTarget {
					continue
				}

				// 重複的將不多次處理
				if _, ok := targetFilters[filter]; ok {
					continue
				}

				targetFilters[filter] = struct{}{}
				filter.Res = append(filter.Res, *token)
			}

		case html.EndTagToken:
			// 處理前一次篩選器
			for filter := range lasttargetFilters {
				// 檢查資料筆數是否相等, 不相等做填補處理
				if len(filter.SubRes) != len(filter.SubContent) {
					filter.SubContent = append(filter.SubContent, "")
				}

				if token.Data == filter.Res[len(filter.Res)-1].Data {
					delete(lasttargetFilters, filter)
				}
			}

			// end 會有額外的空白字串
			tokenizer.Next()

		case html.TextToken:

			// 處理前一次篩選器
			for filter := range lasttargetFilters {
				if (util.FastSearchWithInt(FilterOperation_GetSubcContent, filter.Operation)) != -1 {
					str, _ := util.Big5ToUtf8(tokenizer.Raw())
					filter.SubContent = append(filter.SubContent, str)
				}
			}

			// 取得內文處理
			for filter := range targetFilters {
				if (util.FastSearchWithInt(FilterOperation_GetContent, filter.Operation)) != -1 {
					str, _ := util.Big5ToUtf8(tokenizer.Raw())
					filter.Content = append(filter.Content, str)
				}
			}

			// 排除不取得子物件的篩選器
			for filter := range targetFilters {
				isSubToken := (util.FastSearchWithInt(FilterOperation_GetSubToken, filter.Operation)) != -1
				isSubContent := (util.FastSearchWithInt(FilterOperation_GetSubcContent, filter.Operation)) != -1

				if !isSubToken && !isSubContent {
					delete(targetFilters, filter)
				}
			}

			// 當前篩選器轉移到之前的篩選器
			for filter := range targetFilters {
				if _, ok := lasttargetFilters[filter]; ok {
					continue
				}
				lasttargetFilters[filter] = struct{}{}
			}
			targetFilters = map[*FilterObj]struct{}{}
		}
	}
}

// func sameTargetProcess(tokenizer *html.Tokenizer, tokenType html.TokenType, token *html.Token, targetFilters []*FilterObj) {

// 	switch token.Data {
// 	case "script":

// 		for _, filter := range targetFilters {

// 			// 找到目標 token
// 			filter.Res = append(filter.Res, *token)

// 			// 取得內文
// 			tokenizer.Next()
// 			filter.Content = append(filter.Content, string(tokenizer.Raw()))
// 		}
// 	default:

// 		for _, filter := range targetFilters {
// 			// 找到目標 token
// 			filter.Res = append(filter.Res, *token)
// 		}

// 		tokenizer.Next()
// 		for _, filter := range targetFilters {
// 			isContent := (util.FastSearchWithInt(FilterOperation_GetContent, filter.Operation)) != -1
// 			if isContent {
// 				// 取得內文

// 				for _, filter := range targetFilters {
// 					filter.Content = append(filter.Content, string(tokenizer.Raw()))
// 				}
// 			}
// 		}

// 		// 底下子物件
// 		for {
// 			subtokenType := tokenizer.Next()
// 			subtoken := tokenizer.Token()

// 			// 主標籤結束
// 			if subtokenType == html.EndTagToken && subtoken.Data == token.Data {
// 				break
// 			}

// 			// 排除不是標籤開頭的流程
// 			if subtokenType != html.StartTagToken {
// 				continue
// 			}

// 			for _, filter := range targetFilters {
// 				isSubToken := (util.FastSearchWithInt(FilterOperation_GetSubToken, filter.Operation)) != -1
// 				isSubContent := (util.FastSearchWithInt(FilterOperation_GetSubcContent, filter.Operation)) != -1

// 				if isSubToken {
// 					filter.SubRes = append(filter.SubRes, subtoken)
// 				}

// 				if isSubContent {
// 					tokenizer.Next()
// 					filter.SubContent = append(filter.SubContent, string(tokenizer.Raw()))
// 				}

// 			}
// 		}
// 	}
// }

func AttrCompare(targetAttr html.Attribute, token html.Token) bool {
	for _, attr := range token.Attr {
		if targetAttr == attr {
			return true
		}
	}
	return false
}
