package main

import (
	"testing"
	"yangServer/load/project/goloader"
)

func TestGolandProjectLoad(t *testing.T) {
	goloader.LoadRoot("./")

}
