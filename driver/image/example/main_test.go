package example

import (
	"testing"
	"ycore/driver/image/png"
	"ycore/module/myhtml"
)

func TestDo(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	_, doc, _ := myhtml.GetWebPackage(url)
	png.NewImage(doc, "./", "newimage")
}
