package main

import (
	"log"
	"scraper/internal/cache"
	_dentalstall "scraper/internal/clients/dentalstall"
	"scraper/internal/handler"
	"scraper/internal/repo"
	"scraper/internal/service"
	"scraper/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const TOKEN = "your_static_token"

func main() {
	app := fiber.New()

	app.Use(logger.New())
	url := "https://dentalstall.com/"
	dentaClient := _dentalstall.NewClient(url)
	redis := cache.NewRedisClient()
	cache := cache.NewCache(redis)
	filePath := "./database.json"
	repo := repo.NewRepository(filePath)
	service := service.NewService(dentaClient, cache, repo)
	handler := handler.NewHandler(service)
	app.Post("/scrape", middleware.Authenticate, handler.Scrape)

	log.Fatal(app.Listen(":3000"))
}
