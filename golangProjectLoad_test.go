package main

import (
	"testing"

	"github.com/YWJSonic/ycore/driver/load/project/goloader"
)

func TestGolandProjectLoad(t *testing.T) {
	goloader.LoadRoot("./")

}
