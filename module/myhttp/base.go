package myhttp

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

type MyClient struct {
	dirver *http.Client
}

func New() *MyClient {
	jar, _ := cookiejar.New(nil)
	myHttp := &MyClient{
		dirver: &http.Client{
			Jar: jar,
		},
	}
	return myHttp
}

func (h *MyClient) GetClient() *http.Client {
	return h.dirver
}

//	取得網頁
//	@parame string 網址
//
//	@retrun map[string]string Http Header
//	@return []byte Http Body
//	@return error	錯誤回傳
func (h *MyClient) Get(url url.URL) (map[string][]string, []byte, error) {
	resp, err := h.dirver.Get(url.Path)
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
func (h *MyClient) PostJson(url url.URL, data string) (map[string][]string, []byte, error) {
	resp, err := h.dirver.Post(url.Path, "application/json", strings.NewReader(data))
	if err != nil {
		return nil, nil, err
	}

	return loadHttpRespont(resp)
}
