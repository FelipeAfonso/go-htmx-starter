package home

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HomeHandler(c *fiber.Ctx) error {
	name := "John"
	page := home(name)
	template := templ.Handler(page)
	return adaptor.HTTPHandler(template)(c)
}
