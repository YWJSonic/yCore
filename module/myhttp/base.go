package myhttp

import (
	"log"
	"net/http"
)

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
