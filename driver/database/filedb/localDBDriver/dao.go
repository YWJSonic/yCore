package localDBDriver

type Driver struct {
	path string
}

func NewDriver(setting struct{ Path string }) *Driver {
	return &Driver{
		path: setting.Path,
	}
}

type insertResult struct {
	Key string
}

type fileHead struct {
	Incr   int64            `json:"incr"`
	KeyMap map[string]int64 `json:"keyMap"` // <key, 資料開始位置>
}
type dataHead struct {
	dataStartIndex int64 // 資料起始位置
	dataLenght     int64 // 資料長度
}

func newFileHead() *fileHead {
	return &fileHead{
		KeyMap: map[string]int64{},
	}
}
