package myhttp

import "net/http"

type MyHttp struct {
	client *http.Client
}

func New() *MyHttp {
	myHttp := &MyHttp{
		client: &http.Client{},
	}
	return myHttp
}
