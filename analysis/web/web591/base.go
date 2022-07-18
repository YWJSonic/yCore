package web591

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"yangServer/module/myhtml"
	"yangServer/types"

	"golang.org/x/net/html"
)

type LoginData struct {
	CsrfToken string
	Session   string
	Deviceid  string
	Device    string
}

func GetData(csrfToken string, url string) *HomeList {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-CSRF-TOKEN", csrfToken)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	payload := &HomeList{}
	err = json.Unmarshal(sitemap, payload)
	if err != nil {
		fmt.Println(err)
	}

	return payload
}

func GetDetail(authData LoginData, url string) *HomeDetail {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-CSRF-TOKEN", authData.CsrfToken)
	req.Header.Set("token", authData.Session)
	req.Header.Set("device", authData.Device)
	req.Header.Set("deviceid", authData.Deviceid)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	payload := &HomeDetail{}
	err = json.Unmarshal(sitemap, payload)
	if err != nil {
		fmt.Println(err)
	}
	return payload
}

func Login(device string) LoginData {
	data := LoginData{
		Device: device,
	}
	req, _ := http.NewRequest("GET", LoginPage, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// loginHtml := html.NewTokenizer(res.Body)

	filters := map[types.TokenTypeName]*myhtml.FilterObj{
		"meta": {
			FiltAttrs: []html.Attribute{
				{
					Key: "name",
					Val: "csrf-token",
				},
			},
		},
	}
	// myhtml.HtmlLoopFilterOne(loginHtml, filters)
	for _, htmlToken := range filters["meta"].Res {
		for _, attr := range htmlToken.Attr {
			if attr.Key == "content" {
				data.CsrfToken = attr.Val
			}
		}
	}

	for _, cookie := range res.Cookies() {
		if cookie.Name == "PHPSESSID" {
			data.Session = cookie.Value
		}
		if cookie.Name == "T591_TOKEN" {
			data.Deviceid = cookie.Value
		}
	}

	return data
}
