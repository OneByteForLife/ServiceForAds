package app

import (
	"ServiceForAds/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func SetUpRouters(app *fiber.App) {
	app.Get("/api/v1/", controller.Start)
}

func Run() {
	app := fiber.New()
	app.Use(logger.New())

	SetUpRouters(app)

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("Err up service - %s", err)
	}
}
