package webapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func InitAll(api Handler) {
	app := fiber.New()

	app.Use(logger.New())

	// Регистрация роутеров
	app.Get("/api/v1", Start)
	app.Get("/api/v1/ads/get/all", api.GetAllAds)
	app.Get("/api/v1/ads/get", api.GetAdsByID)
	app.Post("/api/v1/ads/create", api.CreateAds)

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("err up service: %s", err)
	}
}
