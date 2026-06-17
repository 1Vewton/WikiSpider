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
	/*
		settings.project_name = GetEnvString(
			"PROJECT_NAME",
			"WikiSpider",
		)
		settings.version = GetEnvString(
			"VERSION",
			"0.1.0",
		)
		settings.project_url = GetEnvString(
			"PROJECT_URL",
			"https://github.com/1Vewton/WikiSpider",
		)
		settings.request_packge_name = GetEnvString(
			"REQUEST_PACKGE_NAME",
			"Go-http-client/1.1",
		)
		// Log the success message
		service_logger.Info("Configuration initialized successfully...")
	*/
}
