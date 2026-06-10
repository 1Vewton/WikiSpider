package config

import (
	"os"
)

// Get the env variable as a string
func GetEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		// Return the default value if the env variable does not exist
		return defaultValue
	}
	// Return the value if it exists
	return value
}
