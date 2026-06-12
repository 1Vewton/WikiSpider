// Managing the logging system
package logger

import (
	"log/slog"
)

var handler *slog.TextHandler
var logger *slog.Logger
