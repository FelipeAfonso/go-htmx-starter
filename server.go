package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-htmx-starter/internal"
	"go-htmx-starter/pages/home"
	"os"
)

// splitting individual page routing to a separate func for clarity
func router(app *fiber.App) {
	app.Get("/", home.HomeHandler)
	app.Post("/api/increment", home.IncrementCounter)
	// create your handlers and link them in a route here
	app.Get("/", home.HomeHandler)
}

func StartServer() {
	build_id := uuid.New().String()
	app := fiber.New()
	manifest, _ := internal.LoadManifest()
	app.Use(internal.ManifestHandler(manifest))
	app.Use(internal.HotReloadLogger())
	app.Static("/static", "./static")
	app.Static("/dist", "./dist")

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
