package main_test

import (
	"testing"
	"ycore/module/myhtml"
	"ycore/output/image/png"
)

// 從網路抓圖片存到本機
func Test_Image(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	_, doc, _ := myhtml.GetWebPackage(url)
	png.NewImage(doc, "./", "newimage")
}
