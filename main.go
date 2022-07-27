package main

import (
	"flag"
	"ycore/config"
)

var configPath = flag.String("config", "./env.yaml", "specific config to processing")

func main() {
	config.Init(*configPath)

}
