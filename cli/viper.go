package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix = "DPM"
)

// viperSetup binds to the cmd flags, then checks the env, and config file for those values
func viperSetup(cmd *cobra.Command) error {
	// Bing flags to viper
	// Happens first so viper can get the keys to search the other places
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	// Environment Variables
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	// Config
	viper.SetConfigFile(viper.GetString("config"))

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// Config file not found, initialize config file
		viper.SafeWriteConfig()
		// log.Fatal("Please populate config file with appropriate values")
	}
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	return nil
}
