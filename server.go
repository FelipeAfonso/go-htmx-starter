package main

import (
	"github.com/gofiber/fiber/v2"
	"go-htmx-starter/internal"
	"go-htmx-starter/pages/home"
	"os"
)

// splitting individual page routing to a separate func for clarity
func router(app *fiber.App) {
	app.Get("/", home.HomeHandler)
	// create your handlers and link them in a route here
	app.Get("/", home.HomeHandler)
}

func StartServer() {
	app := fiber.New()
	app.Use(internal.HotReloadLogger())
	app.Static("/static", "./static")

	router(app)

	// hot reload stuff
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
