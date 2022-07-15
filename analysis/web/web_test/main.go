package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"yangServer/analysis/web/coolpc"
	"yangServer/analysis/web/web591"
	"yangServer/module/myhtml"
)

var reqDelayTime = time.Second

func main() {
	coolpc.GetWeb()
}

func main_web591() {

	myhtml.SetDefaultClient()

	authData := web591.Login("pc")

	urlStr := web591.Webpage
	payload := web591.GetData(authData.CsrfToken, urlStr)

	lastReqTime := time.Now()
	homeViewMap := map[int64]struct{}{}
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
			lastTime := time.Since(lastReqTime) / time.Second
			lastReqTime = time.Now()
			if lastTime < 2 {
				time.Sleep(time.Second * time.Duration(rand.Int31n(3)+1))
			}

			urlStr = fmt.Sprintf(web591.ObjPage, roomId)
			detailInfo := web591.GetDetail(authData, urlStr)
			idx = strings.Index(detailInfo.Data.FavData.Layout, "衛")
			if idx < 0 {
				continue
			}
			count, _ = strconv.Atoi(detailInfo.Data.FavData.Layout[idx-1 : idx])
			if count < 2 {
				continue
			}
			fmt.Println("targetHouse:", fmt.Sprintf(web591.TargetPage, roomId))
		}

		// 防止過度攻擊
		lastTime := time.Since(lastReqTime) / time.Second
		lastReqTime = time.Now()
		if lastTime < 2 {
			time.Sleep(time.Second * time.Duration(rand.Int31n(3)+1))
		}

		// 下一頁
		count += 30

		// 沒有下一頁
		max, _ := strconv.Atoi(strings.ReplaceAll(payload.Records, ",", ""))
		if count >= max {
			break
		}

		urlStr = fmt.Sprintf(web591.Webpagelast, count, payload.Records)
		payload = web591.GetData(authData.CsrfToken, urlStr)
	}
}
