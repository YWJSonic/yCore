package myhtml

import "golang.org/x/net/html"

type FilterObj struct {
	FiltAttrs  []html.Attribute // 標籤篩選條件
	Res        []html.Token     // 標籤結構
	SubRes     []html.Token     // 標籤結構
	Content    []string         // 標籤內文
	SubContent []string         // 標籤內文
	Operation  []int            // 控制選項 暫時用數字處理
}
