package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Setup custom logger for log can be readable
func Logger() *zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &logger
}
