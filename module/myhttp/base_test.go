package myhttp

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestMyNew(t *testing.T) {
	myhttp := New()

	requestURL := "http://localhost"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := myhttp.client.Do(req)
	fmt.Println(res, err)
	head, b, err := loadHttpRespont(res)
	fmt.Println(head, b, err)
}
