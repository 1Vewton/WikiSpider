package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

var handler *slog.TextHandler
var logger *slog.Logger

type ServiceLogger struct {
	service_name string
}

// initialize logger
func init() {
	opts := &slog.HandlerOptions{AddSource: true}
	_, err := os.Stat("app.log")
	switch {
	case err == nil:
		// Directly open the file
		writer, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		opts.Level = slog.LevelInfo
		handler = slog.NewTextHandler(io.MultiWriter(writer, os.Stdout), opts)
	case os.IsNotExist(err):
		_, err := os.Create("app.log")
		if err != nil {
			panic(err)
		}
		opts.Level = slog.LevelDebug
		writer, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		handler = slog.NewTextHandler(io.MultiWriter(writer, os.Stdout), opts)
	default:
		panic(err)
	}
	logger = slog.New(handler)
}

// Return logger
func NewLogger(service_name string) ServiceLogger {
	var new_logger ServiceLogger
	new_logger.service_name = service_name
	return new_logger
}

// Info
func (s ServiceLogger) Info(msg string) {
	logger.Info(fmt.Sprintf("%s: %s", s.service_name, msg))
}

// Error
func (s ServiceLogger) Error(msg string) {
	logger.Error(fmt.Sprintf("%s: %s", s.service_name, msg))
}
