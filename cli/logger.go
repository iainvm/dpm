package main

import (
	"log/slog"
	"os"

	"github.com/golang-cz/devslog"
	"github.com/spf13/viper"
)

func newLogger() {
	// Default to not logging anything
	var level slog.Level = 9
	if viper.GetBool("verbose") {
		level = slog.LevelDebug
	}

	// Setup a pretty console logger
	logger := slog.New(devslog.NewHandler(os.Stderr, &devslog.Options{
		HandlerOptions: &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		},
	}))

	logger = logger.With(
		slog.String("version", version),
	)
	slog.SetDefault(logger)
}
