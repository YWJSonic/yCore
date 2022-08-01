package main

import (
	"flag"
	"ycore/analysis/web/web591"
)

var configPath = flag.String("config", "./env.yaml", "specific config to processing")

func main() {
	web591.GetData()
}
