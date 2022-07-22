package coolpc

type CacheStruct struct {
	// Key   string `json:"_key"`
	TypeId     int    `json:"typeId"`
	TypeName   string `json:"typeName"`
	Price      int    `json:"price"`      // 價格
	Name       string `json:"name"`       // 標示名稱 未解析
	PriceTag   string `json:"priceTag"`   // 價錢標籤(降價標示)
	UpdateTime int64  `json:"updateTime"` // 更新時間
	Date       string `json:"date"`       // 日期
}
