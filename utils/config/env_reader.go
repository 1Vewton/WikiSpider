package config

import (
	"os"
	"strconv"
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

// Get the env variable as a integer
func GetEnvInteger(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		// Return the default value if the env variable does not exist
		return defaultValue
	}
	result_val, err := strconv.Atoi(value)
	if err != nil {
		// Shut down the process if it is not a int value
		panic(err)
	}
	// Return the result converted to int
	return result_val
}

// Get the env variable as a float64 value
func GetEnvFloat64(key string, defaultValue float64) float64 {
	value := os.Getenv(key)
	if value == "" {
		// Return defalut value if the env variable does not exist
		return defaultValue
	}
	result_var, err := strconv.ParseFloat(value, 64)
	if err != nil {
		// Shut down the process if it is not a float64 value
		panic(err)
	}
	return result_var
}
