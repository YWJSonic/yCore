package myhtml

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

func SetDefaultClient() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	http.DefaultClient.Jar = jar
}
