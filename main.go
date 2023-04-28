package main

import (
	"fmt"
	"github.com/RyaWcksn/nann-e/config"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	"github.com/RyaWcksn/nann-e/server"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	if cfg == nil {
		log.Fatalln("Config is failed to load...")
	}
	fmt.Println(cfg)
	l := logger.New(cfg.App.APPNAME, cfg.App.ENV, cfg.App.LOGLEVEL)
	sv := server.New(cfg, l)

	sv.Start()
}
