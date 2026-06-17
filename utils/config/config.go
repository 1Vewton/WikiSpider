//go:build !test

// Manage the configuration settings
package config

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/utils/logger"
	"github.com/joho/godotenv"
)

var service_logger = logger.NewLogger("Config")

// Initialize the configuration
func init() {
	service_logger.Info("Initializing configuration...")
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(
			"Error loading env file due to %s, using default setting",
			err,
		)
	}
}
