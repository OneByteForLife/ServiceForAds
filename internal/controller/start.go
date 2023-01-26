package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ResponseBody(ver string, status string, code int) fiber.Map {
	return fiber.Map{
		"version": ver,
		"status":  status,
		"code":    code,
	}
}

func Start(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "application/json")

	return c.Status(http.StatusOK).JSON(ResponseBody("1.0.0", "OK", http.StatusOK))
}
