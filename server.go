package main

import (
	"go-htmx-starter/internal"
	"go-htmx-starter/pages/home"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// splitting individual page routing to a separate func for clarity
func router(app *fiber.App) {
	app.Get("/", home.HomeHandler)
	// create your handlers and link them in a route here
	app.Get("/", home.HomeHandler)
}

func StartServer() {
	build_id := uuid.New().String()
	app := fiber.New()
	app.Use(internal.HotReloadLogger())
	app.Static("/static", "./static")

	router(app)

	// hot reload stuff
	isDev := os.Getenv("DEV")
	if isDev == "true" {
		app.Get("/reload", func(c *fiber.Ctx) error {
			c.Status(200)
			return c.SendString(build_id)
		})
	}
	app.Listen(":3000")
}
