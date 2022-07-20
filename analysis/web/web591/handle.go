package web591

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ycore/driver/database/nosql/arango"
	"ycore/module/myhtml"

	"golang.org/x/net/html"
)

// var aMineReqCount = 60
var reqWebCount = 0

func GetData() {

	dbManager, err := arango.New("http://10.146.0.2:8529", "", "", "WebData")
	if err != nil {
		return
	}

	startTime := time.Now()
	fmt.Println("[start] Time:", startTime)
	myhtml.SetDefaultClient()

	authData := login("pc")

	urlStr := Webpage
	payload := getList(authData.CsrfToken, urlStr)

	// lastReqTime := time.Now()
	homeViewMap := map[int64]struct{}{}
	homeResList := []string{}
	count := 0
	idx := 0
	var roomIds []int64
	for {

		// 列出篩選資料 簡易資料
		roomIds = []int64{}
		for _, roomInfo := range payload.Data.Data {

			if _, ok := homeViewMap[roomInfo.PostID]; ok {
				continue
			}

			homeViewMap[roomInfo.PostID] = struct{}{}
			idx = strings.Index(roomInfo.RoomStr, "開放式")
			if idx > 0 {
				continue
			}

			idx = strings.Index(roomInfo.RoomStr, "0廳")
			if idx > 0 {
				continue
			}

			roomIds = append(roomIds, roomInfo.PostID)
		}

		// 查詢詳細資料
		for _, roomId := range roomIds {

			// 防止過度攻擊
			// lastTime := time.Since(lastReqTime) / time.Second
			// lastReqTime = time.Now()
			// if lastTime < 2 {
			// 	time.Sleep(time.Second * time.Duration(rand.Int31n(3)+1))
			// }

			urlStr = fmt.Sprintf(ObjPage, roomId)
			detailInfo := getDetail(authData, urlStr)
			idx = strings.Index(detailInfo.Data.FavData.Layout, "衛")
			if idx < 0 {
				continue
			}
			bscount, _ := strconv.Atoi(detailInfo.Data.FavData.Layout[idx-1 : idx])
			if bscount < 2 {
				continue
			}
			homeResList = append(homeResList, fmt.Sprintf(TargetPage, roomId))
			fmt.Println("[newDetail] Time Spand:", time.Since(startTime), "ReqCount:", reqWebCount)
		}

		// 防止過度攻擊
		// lastTime := time.Since(lastReqTime) / time.Second
		// lastReqTime = time.Now()
		// if lastTime < 2 {
		// 	time.Sleep(time.Second * time.Duration(rand.Int31n(3)+1))
		// }

		// 下一頁
		count += 30

		// 沒有下一頁
		max, _ := strconv.Atoi(strings.ReplaceAll(payload.Records, ",", ""))
		fmt.Println("[nextPage]Now Count", count, "Max Count:", max)
		if count >= max {
			break
		}

		urlStr = fmt.Sprintf(Webpagelast, count, payload.Records)
		payload = getList(authData.CsrfToken, urlStr)
		fmt.Println("[getList] Time Spand:", time.Since(startTime), "ReqCount:", reqWebCount)
	}

	_ = dbManager.Insert(context.TODO(), "FilterHomeData", DBStruct{RoomList: homeResList})
}

func getList(csrfToken string, url string) *HomeList {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-CSRF-TOKEN", csrfToken)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	reqWebCount++

	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	payload := &HomeList{}
	err = json.Unmarshal(sitemap, payload)
	if err != nil {
		fmt.Println("[getList] Unmarshal err:", err)
	}

	return payload
}

func getDetail(authData LoginData, url string) *HomeDetail {
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
	reqWebCount++

	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	payload := &HomeDetail{}
	err = json.Unmarshal(sitemap, payload)
	if err != nil {
		fmt.Println("[getDetail] Unmarshal err:", err)
	}
	return payload
}

func login(device string) LoginData {
	data := LoginData{
		Device: device,
	}
	req, _ := http.NewRequest("GET", LoginPage, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	for _, cookie := range res.Cookies() {
		if cookie.Name == "PHPSESSID" {
			data.Session = cookie.Value
		}
		if cookie.Name == "T591_TOKEN" {
			data.Deviceid = cookie.Value
		}
	}

	loginHtml := html.NewTokenizer(res.Body)

	filters := map[string][]*myhtml.FilterObj{
		"meta": {
			{
				FiltAttrs: []html.Attribute{
					{
						Key: "name",
						Val: "csrf-token",
					},
				},
			},
		},
	}
	myhtml.HtmlLoopFilterOne(loginHtml, filters)

	for _, htmlToken := range filters["meta"][0].Res {
		for _, attr := range htmlToken.Attr {
			if attr.Key == "content" {
				data.CsrfToken = attr.Val
			}
		}
	}

	return data
}
