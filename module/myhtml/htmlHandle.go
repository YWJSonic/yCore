package myhtml

import (
	"io"
	"log"
	"yangServer/dao"
	"yangServer/types"

	"golang.org/x/net/html"
)

func HtmlLoop(tokenizer *html.Tokenizer, filter map[types.TokenTypeName][]*dao.FilterObj) {

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

func HtmlLoopFilterOne(tokenizer *html.Tokenizer, filter map[types.TokenTypeName]*dao.FilterObj) {

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
			if filter, ok := filter[token.Data]; ok {

				isTarget = true // 預設成功

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

func AttrCompare(targetAttr html.Attribute, token html.Token) bool {
	for _, attr := range token.Attr {
		if targetAttr == attr {
			return true
		}
	}
	return false
}
