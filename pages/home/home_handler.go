package home

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go-htmx-starter/pages"
)

func HomeHandler(c *fiber.Ctx) error {
	name := "Mark"
	page := home(name)
	template := templ.Handler(
		pages.Layout(
			"Go + HTMX starter",
			page,
		),
	)
	return adaptor.HTTPHandler(template)(c)
}
