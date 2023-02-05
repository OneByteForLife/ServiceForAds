package webapi

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	BadRequest = "Bad Request"
	Success    = "Success"
)

func ResponseBody(ver string, status string, code int, content interface{}) fiber.Map {
	return fiber.Map{
		"version": ver,
		"status":  status,
		"code":    code,
		"content": content,
	}
}

func Start(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "application/json")

	return c.Status(http.StatusOK).JSON(ResponseBody(
		"1.0.0",
		Success,
		http.StatusOK,
		nil))
}
