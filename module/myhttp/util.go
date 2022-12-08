package myhttp

import (
	"io/ioutil"
	"log"
	"net/http"
)

func SetHeads(req *http.Request, heads map[string]string) {
	for key, value := range heads {
		req.Header.Set(key, value)
	}
}

func GetCookie(req *http.Request) []*http.Cookie {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return res.Cookies()
}

// @params string name
// @params string value
// @params string maxget(s/秒)
//
// @return *http.Cookie
func NewCookie(name, value string, maxget int) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: maxget,
	}
}

func loadHttpRespont(resp *http.Response) (map[string][]string, []byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return resp.Header, nil, err
	}
	return resp.Header, b, nil
}
