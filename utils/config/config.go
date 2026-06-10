package config

import (
	"github.com/1Vewton/WikiSpider/utils/env_reader"
	"github.com/1Vewton/WikiSpider/utils/logger"
	"github.com/joho/godotenv"
)

// Settings struct
type Settings struct {
	wiki_url            string
	email               string
	project_name        string
	project_url         string
	request_packge_name string
	version             string
	maximum_retry       int
}

var settings Settings
var service_logger = logger.NewLogger("Config")

// Initialize the configuration
func init() {
	service_logger.Info("Initializing configuration...")
	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// Fill the .env file
	settings.wiki_url = env_reader.GetEnvString(
		"WIKI_URL",
		"https://en.wikipedia.org/w/api.php?",
	)
	settings.email = env_reader.GetEnvString(
		"EMAIL",
		"example@example.com",
	)
	settings.project_name = env_reader.GetEnvString(
		"PROJECT_NAME",
		"WikiSpider",
	)
	settings.version = env_reader.GetEnvString(
		"VERSION",
		"0.1.0",
	)
	settings.project_url = env_reader.GetEnvString(
		"PROJECT_URL",
		"https://github.com/1Vewton/WikiSpider",
	)
	settings.request_packge_name = env_reader.GetEnvString(
		"REQUEST_PACKGE_NAME",
		"Go-http-client/1.1",
	)
	// Log the success message
	service_logger.Info("Configuration initialized successfully...")
}

// Get wiki url
func GetWikiUrl() string {
	// Return the wiki url
	return settings.wiki_url
}

// Get email
func GetEmail() string {
	// Return the email
	return settings.email
}

// Get project name
func GetProjectName() string {
	// Return the project name
	return settings.project_name
}

// Get project url
func GetProjectUrl() string {
	// Return the project url
	return settings.project_url
}

// Get request package name
func GetRequestPackageName() string {
	// Return the request package name
	return settings.request_packge_name
}

// Get version
func GetVersion() string {
	// Return the version
	return settings.version
}

// Get maximum retry
func GetMaximumRetry() int {
	// Return the maximum retry count
	return settings.maximum_retry
}
