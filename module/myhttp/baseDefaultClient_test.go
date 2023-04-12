package myhttp

import (
	"fmt"
	"testing"
)

func TestMyDefaultHttpGet(t *testing.T) {
	http := NewDefaultClient()
	head, body, err := http.Get("https://www.google.com/manifest?pwa=webhp")
	fmt.Println(head, body, err)
}
