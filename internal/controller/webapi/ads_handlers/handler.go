package adshandler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAdsByID(c *fiber.Ctx) error
	GetAllAds(c *fiber.Ctx) error
	CreateAds(c *fiber.Ctx) error
}
