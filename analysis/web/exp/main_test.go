package main

import (
	"testing"
	"time"
	"ycore/analysis/web/coolpc"
)

var reqDelayTime = time.Second

func Test_main(t *testing.T) {
	coolpc.GetWeb()
}