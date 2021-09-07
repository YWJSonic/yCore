package web591

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	Page1 = "https://rent.591.com.tw/"
	Page2 = "https://www.google-analytics.com/j/collect?v=1&_v=j89&a=1982901210&t=pageview&_s=1&dl=https%3A%2F%2Frent.591.com.tw%2F%3Fkind%3D0%26region%3D1%26rentprice%3D2&ul=zh-tw&de=UTF-8&dt=%E3%80%90%E5%8F%B0%E5%8C%97%E5%B8%82%E5%87%BA%E7%A7%9F%E3%80%91-591%E6%88%BF%E5%B1%8B%E4%BA%A4%E6%98%93%E7%B6%B2&sd=24-bit&sr=1920x1080&vp=1903x362&je=0&_u=aClCAEIhAAAAAG~&jid=1272047502&gjid=651469649&cid=1964670688.1618142734&tid=UA-97423186-1&_gid=1095195379.1618147953&_r=1&gtm=2wg3v0K5L9PD&cd2=gfnhag276rka8m3839shq6psj1&z=1154858152"
	Page3 = "https://union.591.com.tw/stats/event?c=page-pc&a=page-4&l=page-1&_u=q13broe0dcg04oh5fjrh25gom4"
)

type Web591 struct {
	httpClient http.Client
}

func (self *Web591) GetCookie(url *url.URL, name string) *http.Cookie {
	cookies := self.httpClient.Jar.Cookies(url)
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

var request *http.Request

func init() {
	request, _ = http.NewRequest("GET", "https://rent.591.com.tw/home/search/rsList?is_new_list=1&type=1&kind=0&searchtype=1&region=1", nil)
}

func NewWeb591() *Web591 {
	cookie, _ := cookiejar.New(&cookiejar.Options{})
	return &Web591{
		httpClient: http.Client{
			Jar: cookie,
		},
	}
}

func (self *Web591) InitWebData() error {

	resp, err := self.httpClient.Get(Page1)
	if err != nil {
		return err
	}
	self.httpClient.Jar.SetCookies(resp.Request.URL, resp.Cookies())
	println("Page1")
	PrintCookie(self.httpClient.Jar, resp.Request.URL)
	PrintHeader(resp.Header)

	resp, err = self.httpClient.Get(Page2)
	if err != nil {
		return err
	}
	self.httpClient.Jar.SetCookies(resp.Request.URL, resp.Cookies())
	println("Page2")
	PrintCookie(self.httpClient.Jar, resp.Request.URL)
	PrintHeader(resp.Header)

	resp, err = self.httpClient.Get(Page3)
	if err != nil {
		return err
	}
	self.httpClient.Jar.SetCookies(resp.Request.URL, resp.Cookies())
	println("Page3")
	PrintCookie(self.httpClient.Jar, resp.Request.URL)
	PrintHeader(resp.Header)
	return nil
}

func PrintCookie(cookie http.CookieJar, u *url.URL) {
	cookies := cookie.Cookies(u)
	for _, cookie := range cookies {
		fmt.Println(cookie.String())
	}
	fmt.Println("------------------------")
}

func PrintHeader(header http.Header) {
	for k, v := range header {
		fmt.Println(k, v)
	}
	fmt.Println("------------------------------")
}
