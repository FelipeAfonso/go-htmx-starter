package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func HotReloadLogger() fiber.Handler {
	excludeURLs := [2]string{"/reload", "/reload/new"}
	excludeMap := make(map[string]bool)
	for _, url := range excludeURLs {
		excludeMap[url] = true
	}

	defaultLogger := logger.New()

	// Return a custom middleware function
	return func(c *fiber.Ctx) error {
		// Check if the current URL should be excluded
		if _, shouldExclude := excludeMap[c.Path()]; !shouldExclude {
			// If not excluded, use the default logger
			return defaultLogger(c)
		}
		// If excluded, skip logging and continue to the next handler
		return c.Next()
	}
}
