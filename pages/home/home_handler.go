package home

import (
	"go-htmx-starter/pages"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HomeHandler(c *fiber.Ctx) error {
	name := "from the Server!"
	page := home(name)
	template := templ.Handler(
		pages.Layout(
			"Go + HTMX starter",
			page,
		),
	)
	return adaptor.HTTPHandler(template)(c)
}

func IncrementCounter(c *fiber.Ctx) error {
	count := c.FormValue("count", "0")
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return err
	}
	counterTemplate := counter(strconv.Itoa(countInt + 1))
	return adaptor.HTTPHandler(templ.Handler(counterTemplate))(c)
}
