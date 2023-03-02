package main

import (
	"ServiceForAds/config"
	"ServiceForAds/internal/app"
	"ServiceForAds/utils/log"
)

func main() {
	log.ConfigLog()
	conf := config.ReadConfig()

	app.Run(conf)
}
