package config

import (
	"testing"
)

// Test String Env Reading
func TestEnvReadString(t *testing.T) {
	t.Setenv("APP_ENV", "test")
	result := GetEnvString("APP_ENV", "default")
	if result != "test" {
		t.Errorf("Expected 'test', got '%s'", result)
	}
	result2 := "defaults"
	result2 = GetEnvString("APP_DEF", "default")
	if result2 != "default" {
		t.Errorf("Expected 'default', got '%s'", result2)
	}
}

// Test Integer Env Reading
func TestEnvReadInteger(t *testing.T) {
	t.Setenv("APP_ENV_INT", "114514")
	result := GetEnvInteger("APP_ENV_INT", 1919810)
	if result != 114514 {
		t.Errorf("Expected 114514, got %d", result)
	}
	result2 := 114514
	result2 = GetEnvInteger("APP_DEF_INT", 1919810)
	if result2 != 1919810 {
		t.Errorf("Expected 1919810, got %d", result2)
	}
	t.Setenv("APP_ENV_INT", "1.14514")
	result3 := GetEnvFloat64("APP_ENV_INT", 1.919810)
	if result3 != 1.14514 {
		t.Errorf("Expected 1.14514, got %f", result3)
	}
	result4 := 1.14514
	result4 = GetEnvFloat64("APP_DEF_INT", 1.919810)
	if result4 != 1.919810 {
		t.Errorf("Expected 1.919810, got %f", result4)
	}
}
