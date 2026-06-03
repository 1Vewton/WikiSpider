package config

import (
	"github.com/1Vewton/WikiSpider/utils/env_reader"
	"github.com/1Vewton/WikiSpider/utils/logger"
	"github.com/joho/godotenv"
)

// Settings struct
type Settings struct {
	wiki_url string
}

var settings Settings
var service_name string = "Config"

// Initialize the configuration
func init() {
	logger.Info("Initializing configuration...", service_name)
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// Fill the .env file
	settings.wiki_url = env_reader.GetEnvString("WIKI_URL", "https://en.wikipedia.org/w/api.php")
	// Log the success message
	logger.Info("Configuration initialized successfully...", service_name)
}

// Get wiki url
func GetWikiUrl() string {
	return settings.wiki_url
}
