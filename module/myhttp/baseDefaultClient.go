package myhttp

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

// 快速發起請求使用 Default Client
type MyDefaultClient struct{}

func NewDefaultClient() *MyDefaultClient {
	jar, _ := cookiejar.New(nil)
	http.DefaultClient.Jar = jar
	return &MyDefaultClient{}
}

//	取得網頁
//	@parame string 網址
//
//	@retrun map[string]string Http Header
//	@return []byte Http Body
//	@return error	錯誤回傳
func (h *MyDefaultClient) Get(url string) (map[string][]string, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return nil, nil, err
	}

	return loadHttpRespont(resp)
}

//	取得網頁
//	@parame string		網址
//	@parame url.Values	Post資料
//
//	@retrun map[string]string	Http Header
//	@return []byte				Http Body
//	@return error				錯誤回傳
func (h *MyDefaultClient) PostJson(url string, data string) (map[string][]string, []byte, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		return nil, nil, err
	}

	return loadHttpRespont(resp)
}
