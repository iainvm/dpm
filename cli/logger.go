package main

import (
	"log/slog"
	"os"

	"github.com/phsym/console-slog"
	"github.com/spf13/viper"
)

func newLogger() {
	// Default to not logging anything
	var level slog.Level = 9
	if viper.GetBool("verbose") {
		level = slog.LevelDebug
	}

	// Setup a pretty console logger
	logger := slog.New(
		console.NewHandler(
			os.Stderr,
			&console.HandlerOptions{
				Level:     level,
				AddSource: true,
			},
		),
	)
	logger = logger.With(
		slog.String("version", version),
	)
	slog.SetDefault(logger)
}
