package main

import (
	"testing"
	"ycore/load/project/goloader"
)

func TestGolandProjectLoad(t *testing.T) {
	goloader.LoadRoot("./")

}
