package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/iainvm/dpm/dpm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	devDir, err := replaceHome(devDir)
	if err != nil {
		return err
	}

	// Options
	options := &dpm.CloneOptions{}

	// Call dpm actions
	location, err := dpm.Clone(cmd.Context(), devDir, url, options)
	if err != nil {
		return err
	}

	fmt.Fprint(os.Stdout, location)
	return nil
}

func replaceHome(path string) (string, error) {
	if strings.HasPrefix(path, "$HOME") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(homeDir, strings.TrimPrefix(path, "$HOME"))
	}
	return filepath.Abs(path)
}
