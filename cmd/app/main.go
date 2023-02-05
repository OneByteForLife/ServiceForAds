package main

import (
	"ServiceForAds/config"
	"ServiceForAds/internal/app"
	"ServiceForAds/pkg/log"
)

func main() {
	log.ConfigLog()
	conf := config.ReadConfig()

	app.Run(conf)
}
