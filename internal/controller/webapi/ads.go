package webapi

import (
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
		return c.Status(http.StatusBadRequest).JSON(ResponseBody("1.0.0", BadRequest, http.StatusBadRequest, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(ResponseBody(
		"1.0.0",
		Success,
		http.StatusOK,
		ads))
}

func (h *AdsHandler) GetAllAds(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	ads, err := h.service.GetAll(c.Query("limit"), c.Query("offset"), c.Query("sortBy"), c.Query("sortType"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseBody("1.0.0", BadRequest, http.StatusBadRequest, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(ResponseBody(
		"1.0.0",
		Success,
		http.StatusOK,
		ads))
}

func (h *AdsHandler) CreateAds(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	err := h.service.Create(c.Body())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseBody("1.0.0", BadRequest, http.StatusBadRequest, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(ResponseBody(
		"1.0.0",
		Success,
		http.StatusOK,
		nil))
}
