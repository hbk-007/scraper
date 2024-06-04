package handler

import (
	"fmt"
	"scraper/internal/domain/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Scrape(c *fiber.Ctx) error {
	var settings request.ScrapeRequest
	if err := c.BodyParser(&settings); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(settings); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	products, err := h.Service.Scrape(c.Context(), &settings)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	//scraper.SaveToDB(products)
	//scraper.Notify(len(products))

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Scraped %d products successfully.", len(products))})

}
