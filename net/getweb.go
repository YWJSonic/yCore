package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetWebPackage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return nil, err
	}
	return b, nil
}
