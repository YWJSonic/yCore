package png

import (
	"bytes"
	"image/png"
	"os"
	"ycore/util"
)

// 將 memory 資料存到本機 png
func NewImage(data []byte, filePath, fileName string) {
	img, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	f, err := os.Create(util.Sprintf("%s/%s.png", filePath, fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ = png.Encode(f, img)
}
