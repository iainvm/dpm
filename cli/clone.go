package main

import (
	"log/slog"

	"github.com/iainvm/dpm/dpm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cloneCmd parses the CLI args, and calls the dpm.Clone command
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

	// Parse args
	url := args[0]
	projectsDir := viper.GetString("projects-home")

	// Call dpm actions
	err := dpm.Clone(cmd.Context(), projectsDir, url)
	if err != nil {
		return err
	}

	return nil
}
