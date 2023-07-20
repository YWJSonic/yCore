package png

import (
	"bytes"
	"image/png"
	"os"
	"path/filepath"
)

// 將 memory 資料存到本機 png
func NewImage(data []byte, filePath, fileName string) error {
	img, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	path := filepath.Join(filePath, fileName+".png")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = png.Encode(f, img); err != nil {
		return err
	}
	return nil
}
