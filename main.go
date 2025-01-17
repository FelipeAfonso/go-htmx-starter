package main

import (
	"go-htmx-starter/internal"
	"go-htmx-starter/pages/home"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	app := fiber.New()

	app.Use(internal.HotReloadLogger())

	app.Static("/static", "./static")

	app.Get("/", home.HomeHandler)

	isDev := os.Getenv("DEV")

	if isDev == "true" {
		app.Get("/reload", func(c *fiber.Ctx) error {
			key, err := internal.ReadHotReloadKey()
			if err != nil {
				return err
			}
			c.Status(200)
			return c.SendString(key)
		})

		app.Post("/reload/new", func(c *fiber.Ctx) error {
			key, err := internal.WriteHotReloadKey()
			if err != nil {
				return err
			}
			c.Status(201)
			return c.SendString(key)
		})

	}
	app.Listen(":3000")
}
