//go:build test

package logger

import (
	"fmt"
	"log/slog"
	"os"
)

// initialize logger for testing purposes
func init() {
	fmt.Println("Initializing logger for testing purposes")
	handler = slog.NewTextHandler(
		os.Stdout,
		nil,
	)
	logger = slog.New(handler)
}
