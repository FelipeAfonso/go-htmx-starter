package main

import (
	home "go-htmx-starter/pages"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Static("/static", "./static")

	app.Get("/", home.HomeHandler)

	app.Listen(":3000")
}
