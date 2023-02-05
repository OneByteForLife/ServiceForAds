package adshandler

import (
	"ServiceForAds/internal/controller/webapi"
	usecase "ServiceForAds/internal/usecase/ads"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AdsHandler struct {
	service usecase.Service
}

func NewHandler(s usecase.Service) Handler {
	return &AdsHandler{s}
}

func (h *AdsHandler) GetAdsByID(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	ads, err := h.service.GetOne(c.Query("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(webapi.ResponseBody("1.0.0", webapi.BadRequest, http.StatusBadRequest, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(webapi.ResponseBody(
		"1.0.0",
		webapi.Success,
		http.StatusOK,
		ads))
}

func (h *AdsHandler) GetAllAds(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	ads, err := h.service.GetAll(c.Query("limit"), c.Query("offset"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(webapi.ResponseBody("1.0.0", webapi.BadRequest, http.StatusBadRequest, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(webapi.ResponseBody(
		"1.0.0",
		webapi.Success,
		http.StatusOK,
		ads))
}

func (h *AdsHandler) CreateAds(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	err := h.service.Create(c.Body())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(webapi.ResponseBody("1.0.0", webapi.BadRequest, http.StatusBadRequest, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(webapi.ResponseBody(
		"1.0.0",
		webapi.Success,
		http.StatusOK,
		nil))
}
