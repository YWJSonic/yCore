package png

import (
	"bytes"
	"fmt"
	"image/png"
	"os"
)

// 將 memory 資料存到本機 png
func NewImage(data []byte, filePath, fileName string) {
	img, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.png", filePath, fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ = png.Encode(f, img)
}
