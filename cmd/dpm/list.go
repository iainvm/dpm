package main

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/dpm"
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

	// Parse and replace HOME in dev directory path
	devDir := viper.GetString("dev-directory")

	// Options
	options := &dpm.ListOptions{
		Names: viper.GetBool("names"),
	}

	list, err := dpm.List(cmd.Context(), devDir, options)
	if err != nil {
		return err
	}

	for _, project := range list {
		_, err = fmt.Fprint(output, project, "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
