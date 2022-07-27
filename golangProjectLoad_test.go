package main

import (
	"testing"
	"ycore/driver/load/project/goloader"
)

func TestGolandProjectLoad(t *testing.T) {
	goloader.LoadRoot("./")

}
