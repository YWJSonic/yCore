package main

import (
	"flag"
	"ycore/worker/websiteapi"
)

var configPath = flag.String("config", "./env.yaml", "specific config to processing")

func main() {
	websiteapi.New()
}
