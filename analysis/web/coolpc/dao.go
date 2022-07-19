package coolpc

type CacheStruct struct {
	// Key   string `json:"_key"`
	TypeId     int    `json:"typeId"`
	TypeName   string `json:"typeName"`
	Price      int    `json:"price"`      // 價格
	Name       string `json:"name"`       // 標示名稱 未解析
	UpdateTime int64  `json:"updateTime"` // 更新時間
}
