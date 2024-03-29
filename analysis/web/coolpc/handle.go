package coolpc

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/YWJSonic/ycore/module/mydb"
	"github.com/YWJSonic/ycore/module/myhtml"
	"github.com/YWJSonic/ycore/module/mylog"
	"github.com/YWJSonic/ycore/types"
	"github.com/YWJSonic/ycore/util"

	"golang.org/x/net/html"
)

func GetWeb() {

	coolpcDB := coolpcDbManager{}
	dbManager, err := mydb.NewArangoDB("http://10.140.0.2:8529", "", "", "WebData")
	if err != nil {
		return
	}
	coolpcDB.Manager = dbManager

	req, _ := http.NewRequest("GET", WebPage, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// b, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(b)
	// return

	// reader := bytes.NewReader(body)
	loginHtml := html.NewTokenizer(res.Body)

	filters := map[types.TokenTypeName][]*myhtml.FilterObj{}
	for i, count := 1, 31; i < count; i++ {
		filter := &myhtml.FilterObj{}
		filter.FiltAttrs = append(filter.FiltAttrs,
			html.Attribute{
				Key: "name",
				Val: util.Sprintf("n%d", i),
			})
		filter.Operation = append(filter.Operation, myhtml.FilterOperation_GetSubToken, myhtml.FilterOperation_GetSubcContent)
		filters["select"] = append(filters["select"], filter)
	}
	filters["script"] = append(filters["script"], &myhtml.FilterObj{
		Operation: []int{
			myhtml.FilterOperation_GetContent,
		},
	})
	myhtml.HtmlLoopFilterOne(loginHtml, filters)

	var dataMap map[string]interface{}
	if len(filters["script"][0].Content) > 0 {
		data := filters["script"][0].Content[0]
		dataMap = spliteScriptData(data)
	}

	datetime := util.ParseJavaUnixSec(util.ServerTimeNow())
	date := util.ServerTimeNow().Format("2006-01-02")
	collectionName := "Coolpc"
	wg := sync.WaitGroup{}
	wg.Add(len(filters["select"]))

	for idx, filter := range filters["select"] {
		idx++
		go func(idx int, filter *myhtml.FilterObj) {
			key := util.Sprintf("c%d", idx)
			typeName := typeMap[idx]
			data := dataMap[key].([]int)
			for subidx, token := range filter.SubRes {
				if token.Data == "option" {
					for _, attr := range token.Attr {
						if attr.Key == "value" {
							v, _ := strconv.Atoi(attr.Val)
							price := data[v]

							// 有價格的才是實際商品
							if price > 0 {
								nameSplite := strings.Split(filter.SubContent[subidx], ",")
								if len(nameSplite) >= 2 {
									_ = coolpcDB.Create(context.TODO(), collectionName,
										CacheStruct{
											Date:       date,
											UpdateTime: datetime,
											TypeName:   typeName,
											TypeId:     idx,
											Price:      data[v],
											PriceTag:   nameSplite[len(nameSplite)-1],
											Name:       nameSplite[0],
											OriginName: filter.SubContent[subidx],
										})
								} else {
									_ = coolpcDB.Create(context.TODO(), collectionName,
										CacheStruct{
											Date:       date,
											UpdateTime: datetime,
											TypeName:   typeName,
											TypeId:     idx,
											Price:      data[v],
											Name:       nameSplite[0],
											OriginName: filter.SubContent[subidx],
										})
								}
							}
						}
					}
				}
			}
			wg.Done()
		}(idx, filter)
	}

	wg.Wait()
}

func UpdateOldFrame() {
	coolpcDB := &coolpcDbManager{}
	dbManager, err := mydb.NewArangoDB("http://10.146.0.2:8529", "", "", "WebData")
	if err != nil {
		return
	}
	coolpcDB.Manager = dbManager
	docs, err := coolpcDB.GetAllData(context.TODO(), "Coolpc")

	for _, doc := range docs {
		updatedata(coolpcDB, doc)
	}

}

func updatedata(coolpcDB *coolpcDbManager, data *CacheStruct) {
	if data.PriceTag != "" {
		return
	}

	nameSplite := strings.Split(data.Name, ",")
	if len(nameSplite) >= 2 {
		data.PriceTag = nameSplite[1]
		data.Name = nameSplite[0]
		err := coolpcDB.Update(context.TODO(), "Coolpc", data.Key, data)
		if err != nil {
			mylog.Error(err)
		}
	}
}

func spliteScriptData(data string) map[string]interface{} {
	res := map[string]interface{}{}

	spliteStr := strings.Split(data, "\n")
	for _, dataStr := range spliteStr {
		// 判斷格式
		idx := strings.IndexByte(dataStr, '=')
		if idx < 0 {
			continue
		}

		// 確認資料結構
		key := dataStr[0:idx]

		// 排除 c類型以外資料
		if key[0] != 'c' {
			continue
		}

		valStr := dataStr[idx+1:]
		if len(valStr) < 2 || (valStr[0] != '[' && valStr[len(valStr)-1] != ']') {
			continue
		}

		// 資料解析
		values := []int{}
		valueSplite := strings.Split(dataStr[idx+2:len(dataStr)-1], ",")
		for _, valueStr := range valueSplite {
			val, err := strconv.Atoi(valueStr)
			if err != nil {
				val = -1
			}
			values = append(values, val)
		}

		res[key] = values
	}
	return res
}
