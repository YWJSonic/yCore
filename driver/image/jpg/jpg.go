package jpg

import (
	"bytes"
	"image/jpeg"
	"os"
	"path/filepath"
)

func NewImage(data []byte, filePath, fileName string) error {
	img, err := jpeg.Decode(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	path := filepath.Join(filePath, fileName+".jpeg")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = jpeg.Encode(f, img, nil); err != nil {
		return err
	}
	return nil
}
