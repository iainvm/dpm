package logger

import (
	"log/slog"
	"os"
)

func New(verbose bool) *slog.Logger {
	addSource := false
	logLevel := slog.LevelInfo
	if verbose {
		addSource = true
		logLevel = slog.LevelDebug
	}

	slog.Info(
		"Log level",
		slog.Any("level", logLevel),
	)
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: addSource,
			Level:     slog.LevelDebug,
		},
	)
	// handler := NewDefault(&slog.HandlerOptions{
	// 	AddSource: true,
	// 	Level:     slog.LevelDebug,
	// })

	log := slog.New(handler)

	slog.SetDefault(log)

	return log
}
