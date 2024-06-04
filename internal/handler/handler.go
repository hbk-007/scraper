package handler

import (
	"scraper/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service service.IClient
}

type IHandler interface {
	Scrape(c *fiber.Ctx) error
}

func NewHandler(service service.IClient) Handler {
	return Handler{Service: service}
}
