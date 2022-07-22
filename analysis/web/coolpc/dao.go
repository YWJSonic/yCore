package coolpc

import (
	"context"
	"ycore/module/mydb"

	"github.com/arangodb/go-driver"
)

type CacheStruct struct {
	Key        string `json:"_key,omitempty"`
	TypeId     int    `json:"typeId"`
	TypeName   string `json:"typeName"`
	Price      int    `json:"price"`      // 價格
	Name       string `json:"name"`       // 標示名稱 未解析
	PriceTag   string `json:"priceTag"`   // 價錢標籤(降價標示)
	UpdateTime int64  `json:"updateTime"` // 更新時間
	Date       string `json:"date"`       // 日期
}

type coolpcDbManager struct {
	*mydb.Manager
}

func (self *coolpcDbManager) Update(ctx context.Context, collection string, key string, value interface{}) error {
	_, err := self.Client.Query(ctx,
		`FOR doc IN @@collection
		FILTER doc._key == @key
		UPDATE doc WITH @value in @@collection`,
		map[string]interface{}{
			"@collection": collection,
			"key":         key,
			"value":       value,
		})
	return err
}

func (self *coolpcDbManager) GetAllData(ctx context.Context, collection string) ([]*CacheStruct, error) {
	cursor, err := self.Client.Query(ctx,
		`FOR doc IN @@collection
		return doc`,
		map[string]interface{}{
			"@collection": collection,
		})

	if err != nil {
		return nil, err
	}

	res := []*CacheStruct{}
	for cursor.HasMore() {
		doc := &CacheStruct{}
		_, err = cursor.ReadDocument(ctx, doc)
		if driver.IsNoMoreDocuments(err) {
			break
		}
		res = append(res, doc)
	}
	return res, nil
}
