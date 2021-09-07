package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//	取得網頁
//	@parame string 網址
//	@retrun map[string]string Http Header
//	@return []byte Http Body
//	@return error	錯誤回傳
func GetWebPackage(url string) (map[string][]string, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return nil, nil, err
	}

	return LoadHttpRespont(resp)
}

func LoadHttpRespont(resp *http.Response) (map[string][]string, []byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %v\n", err)
		return resp.Header, nil, err
	}
	return resp.Header, b, nil
}
