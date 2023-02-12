package app

import (
	"ServiceForAds/config"
	"ServiceForAds/internal/controller/webapi"

	usecase "ServiceForAds/internal/usecase/ads"
	"ServiceForAds/pkg/database"
)

func Run(conf *config.Config) {
	pool := database.ConnectToDatabase(conf)

	api := webapi.NewHandler(usecase.NewService(usecase.NewStorage(pool)))

	webapi.InitAll(api)
}
