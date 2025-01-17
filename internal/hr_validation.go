package internal

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

func WriteHotReloadKey() (string, error) {
	id := uuid.New().String()
	// Create or open the .hrid file for writing
	file, err := os.OpenFile(".hrid", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to open .hrid for writing: %w", err)
	}
	defer file.Close()
	// Write the ID to the file
	if _, err := file.WriteString(id); err != nil {
		return "", fmt.Errorf("failed to write unique ID to .hrid: %w", err)
	}
	return id, nil
}

func ReadHotReloadKey() (string, error) {
	data, err := os.ReadFile(".hrid")
	if err != nil {
		return "", fmt.Errorf("failed to read unique ID from .hrid: %w", err)
	}
	return string(data), nil
}
