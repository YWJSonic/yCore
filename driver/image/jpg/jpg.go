package jpg

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"os"
)

func NewImage(data []byte, filePath, fileName string) {
	img, err := jpeg.Decode(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.png", filePath, fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_ = jpeg.Encode(f, img, nil)
}
