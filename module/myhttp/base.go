package myhttp

import "net/http"

type MyClient struct {
	dirver *http.Client
}

func New() *MyClient {
	myHttp := &MyClient{
		dirver: &http.Client{},
	}
	return myHttp
}
