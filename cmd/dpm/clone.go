package main

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/dpm"
)

// cloneCmd parses the CLI args, and calls the dpm.Clone command
func cloneCmd(cmd *cobra.Command, args []string) error {
	slog.DebugContext(
		cmd.Context(),
		"Clone command executed",
		slog.Any("flags", viper.AllSettings()),
		slog.String("used_config", viper.ConfigFileUsed()),
	)

	// Parse args
	url := args[0]

	// Parse and replace HOME in dev directory path
	devDir := viper.GetString("dev-directory")

	// Options
	options := &dpm.CloneOptions{}

	// Call dpm actions
	location, err := dpm.Clone(cmd.Context(), devDir, url, options)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(output, location)
	if err != nil {
		return err
	}

	return nil
}
