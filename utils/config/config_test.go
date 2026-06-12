package config

import (
	"os"
	"testing"
)

// Test env reading
func TestEnvRead(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	result := GetEnvString("APP_ENV", "default")
	if result != "test" {
		t.Errorf("Expected 'test', got '%s'", result)
	}
	result2 := "defaults"
	result2 = GetEnvString("APP_DEF", "default")
	if result2 != "default" {
		t.Errorf("Expected 'default', got '%s'", result)
	}
}
