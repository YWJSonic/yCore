package main

import (
	"flag"
	"ycore/worker/googlelogin"
)

var configPath = flag.String("config", "./env.yaml", "specific config to processing")

func main() {
	googlelogin.New()
}
