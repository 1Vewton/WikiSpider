package logger

import (
	"fmt"
)

// Data structure for ServiceLogger
type ServiceLogger struct {
	service_name string
}

// Return logger
func NewLogger(service_name string) ServiceLogger {
	var new_logger ServiceLogger
	new_logger.service_name = service_name
	return new_logger
}

// Info
func (s ServiceLogger) Info(msg string) {
	if logger != nil {
		logger.Info(fmt.Sprintf("%s: %s", s.service_name, msg))
	}
}

// Error
func (s ServiceLogger) Error(msg string) {
	if logger != nil {
		logger.Error(fmt.Sprintf("%s: %s", s.service_name, msg))
	}
}
