package env_reader

import (
	"os"
)

func GetEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return key
}
