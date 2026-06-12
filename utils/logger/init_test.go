//go:build test

package logger

import (
	"io"
	"log/slog"
	"os"
)

// initialize logger for testing purposes
func init() {
	opts := &slog.HandlerOptions{AddSource: true}
	handler = slog.NewTextHandler(
		io.Writer(os.Stdout),
		opts,
	)
	logger = slog.New(handler)
}
