package web591

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"yangServer/net"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func TestWeb(t *testing.T) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	resp, _ := http.Get(Page1)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	ts := httptest.NewServer(writeHTML(string(body)))
	defer ts.Close()

	chromedp.ListenTarget(ctx,
		func(ev interface{}) {
			if ev, ok := ev.(*network.EventResponseReceived); ok {
				// fmt.Println("event received:")
				// fmt.Println(ev.Type)

				if ev.Type != "XHR" {
					return
				}

				// fmt.Println("= " + ev.Response.RequestHeadersText + " =")

				go func() {
					// print response body
					c := chromedp.FromContext(ctx)
					rbp := network.GetResponseBody(ev.RequestID)
					body, err := rbp.Do(cdp.WithExecutor(ctx, c.Target))
					if err != nil {
						fmt.Println(err)
					}
					if err = ioutil.WriteFile(ev.RequestID.String(), body, 0644); err != nil {
						log.Fatal(err)
					}
					if err == nil {
						fmt.Printf("%s\n", body)
					}
				}()

			}
		},
	)

	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
	); err != nil {
		log.Fatal(err)
	}
}
func TestDo(t *testing.T) {
	Client := NewWeb591()
	// header, doc, err := net.GetWebPackage(Url)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for k, v := range header {
	// 	fmt.Println(k, v)
	// 	if k == "Set-Cookie" {
	// 		Client.Cookie = SplitCookieData(v)
	// 		Client.httpClient.Jar.SetCookies()
	// 	}
	// }
	// fmt.Println(string(doc))

	Client.InitWebData()

	cookies := fmt.Sprintf("T591_TOKEN=%v ", Client.GetCookie(request.URL, "PHPSESSID").Value)
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("X-CSRF-TOKEN", "32eD2FIhr3wAdyxIVUi8SysDrCf3ugPYVBWAVuXT")
	request.Header.Set("Cookie", cookies)
	Client.httpClient.Jar.SetCookies(request.URL, request.Cookies())
	// request.Header.Set("Cookie", "XSRF-TOKEN=eyJpdiI6IjZ2WFp3XC9Femlnc0ZsYkVDYmx6c3JRPT0iLCJ2YWx1ZSI6IjdlbW1SSjFpTFJqbFwvRXlXSHZJWG9wcmZJUXFvb1c3bEVSK29OZHdRdXMrUXlsblhBT01rRWpBS1JVQWFOeW5PYW14ZXBQQnVtbGpCaklveThQbU5ZZz09IiwibWFjIjoiYTMxZjEyNGM5ZGQwZWU2NjkxZGNkOWFhMzZlNWM5YWYwYjU3MGIwMjExZTljMzFkODIwNzU2MDc1OTc4OGU4NSJ9; webp=1; urlJumpIp=1; urlJumpIpByTxt=%E5%8F%B0%E5%8C%97%E5%B8%82; 591_new_session=eyJpdiI6IktQVUMrdGxKdWhyWWRQZERqeld6Snc9PSIsInZhbHVlIjoiTTZyUTd1eTJkYm9yeEpDYjhJQ2Y1T1BLRituMVoxQXViR0dGM0tsZlpRcnNGUlI0c2xqd2NZQk13cERad2xUYzlOOGo5U3lNWWlaTHc2aG1UdWtWamc9PSIsIm1hYyI6IjBlYTkxMTE1NmI2MjkzZGQxZmM1ZWMzOTM0ZDgwY2NlNjliMzQ4ODZiMWFiOWNmY2MyZDJjMjc0MjM2ZDRmMjQifQ%3D%3D; _fbp=fb.2.1618144328831.1217817029; _ga=GA1.4.1964670688.1618142734; _gid=GA1.4.1891432225.1618144329; _gat_UA-97423186-1=1")

	resp, err := Client.httpClient.Do(request)
	if err != nil {
		fmt.Println("Error:", err)
	}
	println("Page4")
	PrintHeader(resp.Header)
	PrintCookie(Client.httpClient.Jar, resp.Request.URL)
}

func TestDo2(t *testing.T) {

	req, _ := http.NewRequest("GET", "https://rent.591.com.tw/home/search/rsList?is_new_list=1&type=1&kind=1&searchtype=1&region=1", nil)

	req.Header.Set("Accept-Encoding", " gzip, deflate, br")
	req.Header.Set("Cookie", " webp=1; PHPSESSID=duvsap7gpp54ejrir47l37iir0; urlJumpIp=1; urlJumpIpByTxt=%E5%8F%B0%E5%8C%97%E5%B8%82; T591_TOKEN=duvsap7gpp54ejrir47l37iir0; new_rent_list_kind_test=0; 591_new_session=eyJpdiI6InNFa0w5cVE3QU56RW43NnBjWUlEeGc9PSIsInZhbHVlIjoiMnhETmxVU1ZKZmJFNUtxaWhRNnVzTlFEelwvXC9aeDM5NFR1anZnWHlWNHVTcFNmaCtIXC9IWWxpMWdObXBzTzFDTXpMV21jNkV4dWwwUVE5blVJUU1nRnc9PSIsIm1hYyI6IjEzNDQ1YmI5ZDJjYTlmOWFmM2VmYTVmMDI5OTkyMjNhY2MzNmRkZGY4MWMyMzg4Y2IzMWU4YTcxMGY1ZTM1NzkifQ%3D%3D; _ga=GA1.3.1092790415.1618158016; _gid=GA1.3.1112558231.1618158016; _gat=1; _ga=GA1.4.1092790415.1618158016; _gid=GA1.4.1112558231.1618158016; _dc_gtm_UA-97423186-1=1")
	req.Header.Set("X-CSRF-TOKEN", " C3Tr9ojctqEVfHfGJU6tGc03WQBPXiAP7Av1UjSu")

	res, err := http.DefaultClient.Do(req)
	head, body, err := net.LoadHttpRespont(res)
	// data := map[string]interface{}{}
	// err = json.Unmarshal(body, &data)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	zipFile, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	FullData, _ := ioutil.ReadAll(zipFile)

	fmt.Println(head, err)
	fmt.Println(string(FullData))
}

// Header
// **
// Accept-Encoding: gzip, deflate, br
// Cookie: webp=1; PHPSESSID=duvsap7gpp54ejrir47l37iir0; urlJumpIp=1; urlJumpIpByTxt=%E5%8F%B0%E5%8C%97%E5%B8%82; T591_TOKEN=duvsap7gpp54ejrir47l37iir0; new_rent_list_kind_test=0; 591_new_session=eyJpdiI6InNFa0w5cVE3QU56RW43NnBjWUlEeGc9PSIsInZhbHVlIjoiMnhETmxVU1ZKZmJFNUtxaWhRNnVzTlFEelwvXC9aeDM5NFR1anZnWHlWNHVTcFNmaCtIXC9IWWxpMWdObXBzTzFDTXpMV21jNkV4dWwwUVE5blVJUU1nRnc9PSIsIm1hYyI6IjEzNDQ1YmI5ZDJjYTlmOWFmM2VmYTVmMDI5OTkyMjNhY2MzNmRkZGY4MWMyMzg4Y2IzMWU4YTcxMGY1ZTM1NzkifQ%3D%3D; _ga=GA1.3.1092790415.1618158016; _gid=GA1.3.1112558231.1618158016; _gat=1; _ga=GA1.4.1092790415.1618158016; _gid=GA1.4.1112558231.1618158016; _dc_gtm_UA-97423186-1=1
// X-CSRF-TOKEN: C3Tr9ojctqEVfHfGJU6tGc03WQBPXiAP7Av1UjSu
// X-Requested-With: XMLHttpRequest
// Accept: application/json, text/javascript, */*; q=0.01
// Accept-Language: zh-TW,zh;q=0.9
// Connection: keep-alive
// Host: rent.591.com.tw
// Referer: https://rent.591.com.tw/?kind=0&region=1&rentprice=2
// sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"
// sec-ch-ua-mobile: ?0
// Sec-Fetch-Dest: empty
// Sec-Fetch-Mode: cors
// Sec-Fetch-Site: same-origin
// User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36
// **

// PHPSESSID=hjbhphhqpbksqn4skj18leiss3
// urlJumpIp=1
// urlJumpIpByTxt=%E5%8F%B0%E5%8C%97%E5%B8%82
// new_rent_list_kind_test=1
// 591_new_session=eyJpdiI6Im9mTGYxOHg1ODNOYmg0YVB5anN2aHc9PSIsInZhbHVlIjoicmVpVmlkTE53dVwvN3RWWnQ1OTFSWWVZWDNpZnhNZ2YwM1ZWR1ZzaVpCemlHUzdSVG1ZQXdiczQrTFBhakVyVW5zOVJPZ21GUnhaa2FoU0ZzRGVXS3JRPT0iLCJtYWMiOiIyMjM2ODE4MWVlMWE3ZTUzYmQyMTlkYTgwZTg2Nzg5MjhiMmM0YjIwNzc5ODgzYWI1NGVlYjE1NzEwNGRlMWEyIn0%3D

// webp=1;
// PHPSESSID=duvsap7gpp54ejrir47l37iir0;
// urlJumpIp=1;
// urlJumpIpByTxt=%E5%8F%B0%E5%8C%97%E5%B8%82;
// T591_TOKEN=duvsap7gpp54ejrir47l37iir0;
// new_rent_list_kind_test=0;
// 591_new_session=eyJpdiI6InNFa0w5cVE3QU56RW43NnBjWUlEeGc9PSIsInZhbHVlIjoiMnhETmxVU1ZKZmJFNUtxaWhRNnVzTlFEelwvXC9aeDM5NFR1anZnWHlWNHVTcFNmaCtIXC9IWWxpMWdObXBzTzFDTXpMV21jNkV4dWwwUVE5blVJUU1nRnc9PSIsIm1hYyI6IjEzNDQ1YmI5ZDJjYTlmOWFmM2VmYTVmMDI5OTkyMjNhY2MzNmRkZGY4MWMyMzg4Y2IzMWU4YTcxMGY1ZTM1NzkifQ%3D%3D;
// _ga=GA1.3.1092790415.1618158016;
// _gid=GA1.3.1112558231.1618158016;
// _gat=1;
// _ga=GA1.4.1092790415.1618158016;
// _gid=GA1.4.1112558231.1618158016;
// _dc_gtm_UA-97423186-1=1
