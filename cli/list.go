package main

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func listCmd(cmd *cobra.Command, args []string) error {
	slog.DebugContext(
		cmd.Context(),
		"List command executed",
		slog.Group(
			"flags",
			slog.String("config", viper.GetString("config")),
			slog.Bool("verbose", viper.GetBool("verbose")),
			slog.Bool("names", viper.GetBool("names")),
		),
		slog.String("used_config", viper.ConfigFileUsed()),
	)

	return nil
}
