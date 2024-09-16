package logger

import (
	"log/slog"
	"os"
)

func MustSetupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case "env":
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		)
	case "dev":
		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout, &slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			),
		)
	case "prod":
		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		logger = slog.New(
			slog.NewJSONHandler(
				f, &slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		)
	}

	return logger
}
