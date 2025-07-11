package main

import (
	"log/slog"

	"github.com/iainvm/dpm/dpm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func cloneCmd(cmd *cobra.Command, args []string) error {
	slog.DebugContext(
		cmd.Context(),
		"Clone command executed",
		slog.Group(
			"flags",
			slog.String("config", viper.GetString("config")),
			slog.String("projects-home", viper.GetString("projects-home")),
			slog.Bool("verbose", viper.GetBool("verbose")),
			slog.Bool("short", viper.GetBool("short")),
			slog.String("identity-file", viper.GetString("identity-file")),
		),
		slog.String("used_config", viper.ConfigFileUsed()),
	)
	// parse args
	url := args[0]
	err := dpm.Clone(cmd.Context(), viper.GetString("projects-home"), url)
	if err != nil {
		return err
	}

	return nil
}
