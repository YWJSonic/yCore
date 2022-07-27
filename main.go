package main

import (
	"flag"
	"fmt"
	"ycore/config"
)

var configPath = flag.String("config", "./env.yaml", "specific config to processing")

func main() {
	config.Init(*configPath)
	fmt.Println(config.EnvInfo)
}
