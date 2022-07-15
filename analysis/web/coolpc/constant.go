package coolpc

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var WebPage = "https://www.coolpc.com.tw/evaluate.php"

func GetWeb() {

	req, _ := http.NewRequest("GET", WebPage, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

	// loginHtml := html.NewTokenizer(res.Body)

	// filters := map[types.TokenTypeName]*dao.FilterObj{
	// 	"meta": {
	// 		FiltAttrs: []html.Attribute{
	// 			{
	// 				Key: "name",
	// 				Val: "csrf-token",
	// 			},
	// 		},
	// 	},
	// }
	// myhtml.HtmlLoopFilterOne(loginHtml, filters)
	// for _, htmlToken := range filters["meta"].Res {
	// 	for _, attr := range htmlToken.Attr {
	// 		if attr.Key == "content" {
	// 			data.CsrfToken = attr.Val
	// 		}
	// 	}
	// }

	// for _, cookie := range res.Cookies() {
	// 	if cookie.Name == "PHPSESSID" {
	// 		data.Session = cookie.Value
	// 	}
	// 	if cookie.Name == "T591_TOKEN" {
	// 		data.Deviceid = cookie.Value
	// 	}
	// }

	// return data
}
