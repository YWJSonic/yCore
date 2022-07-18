package util

import (
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Big5ToUtf8(source []byte) (string, error) {
	big5Toutf8 := traditionalchinese.Big5.NewDecoder()
	str, _, err := transform.String(big5Toutf8, string(source))
	return str, err
}
