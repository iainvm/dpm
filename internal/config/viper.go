package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Load loads the configuration from the default file or the given configFile path
func Load(prefix string, configPath string) error {
	slog.Debug(
		"Loading config file",
		slog.String("file", viper.ConfigFileUsed()),
	)

	if configPath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configPath)
	} else {
		// Find home directory.
		userHome, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in user's home directory with name ".dpm" (without extension).
		viper.AddConfigPath(fmt.Sprintf("%s/.config/dpm", userHome))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		slog.Error(
			"Loaded config file",
			slog.String("file", viper.ConfigFileUsed()),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}
