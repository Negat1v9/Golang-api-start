package logger

import (
	"log/slog"
	"os"
)

// creating the default logger for the service
func InitLogger() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(log)

	log.Info("set default logger, logger level: [DEBUG]")
}
