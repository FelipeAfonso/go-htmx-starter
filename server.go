package main

import (
	"github.com/gofiber/fiber/v2"
	"go-htmx-starter/internal"
	"go-htmx-starter/pages/home"
)

// splitting individual page routing to a separate func for clarity
func router(app *fiber.App) {
	app.Get("/", home.HomeHandler)
	app.Post("/api/increment", home.IncrementCounter)
	// create your handlers and link them in a route here
	app.Get("/", home.HomeHandler)
}

func StartServer() {
	app := fiber.New()
	manifest, _ := internal.LoadManifest()
	app.Use(internal.ManifestHandler(manifest))
	app.Use(internal.HotReloadLogger())
	app.Static("/static", "./static")
	app.Static("/dist", "./dist")
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	router(app)
	app.Listen(":3000")
}
