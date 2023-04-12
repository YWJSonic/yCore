package jpg

import (
	"bytes"
	"image/jpeg"
	"os"

	"github.com/YWJSonic/ycore/util"
)

func NewImage(data []byte, filePath, fileName string) {
	img, err := jpeg.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	f, err := os.Create(util.Sprintf("%s/%s.png", filePath, fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_ = jpeg.Encode(f, img, nil)
}
