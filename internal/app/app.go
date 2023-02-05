package app

import (
	"ServiceForAds/config"
	"ServiceForAds/internal/controller/webapi"
	adshandler "ServiceForAds/internal/controller/webapi/ads_handlers"
	usecase "ServiceForAds/internal/usecase/ads"
	"ServiceForAds/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func Run(conf *config.Config) {
	pool, err := database.ConnectToDatabase(conf)
	if err != nil {
		logrus.Errorf("error connect to database: %s", err)
	}

	api := adshandler.NewHandler(usecase.NewService(usecase.NewStorage(pool)))

	app := fiber.New()

	app.Use(logger.New())

	// Регистрация роутеров
	app.Get("/api/v1", webapi.Start)
	app.Get("/api/v1/ads/get/all", api.GetAllAds)
	app.Get("/api/v1/ads/get", api.GetAdsByID)
	app.Post("/api/v1/ads/create", api.CreateAds)

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("err up service: %s", err)
	}
}
