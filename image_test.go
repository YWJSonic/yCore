package main_test

import (
	"testing"
	"yangServer/net"
	"yangServer/output/image/png"
)

// 從網路抓圖片存到本機
func Test_Image(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	doc, _ := net.GetWebPackage(url)
	png.NewImage(doc, "./", "newimage")
}
