package internal

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"os"
)

// ManifestEntry represents an entry in the Vite manifest
type ManifestEntry struct {
	File   string   `json:"file"`
	Src    string   `json:"src"`
	CSS    []string `json:"css,omitempty"`
	Assets []string `json:"assets,omitempty"`
}

// LoadManifest loads the Vite manifest file
func LoadManifest() (map[string]ManifestEntry, error) {
	fallback := make(map[string]ManifestEntry)
	data, err := os.ReadFile("dist/.vite/manifest.json")
	if err != nil {
		return fallback, err
	}

	var manifest map[string]ManifestEntry
	if err := json.Unmarshal(data, &manifest); err != nil {
		return fallback, err
	}

	return manifest, nil
}

func ManifestHandler(manifest map[string]ManifestEntry) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("manifest", manifest)
		c.Locals("isDev", os.Getenv("GO_ENV") != "production")
		return c.Next()
	}
}
