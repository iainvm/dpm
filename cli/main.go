package main

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/charmbracelet/fang"
	"github.com/phsym/console-slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type command struct {
	command     *cobra.Command
	flags       func(cmd *cobra.Command)
	subcommands []command
}

const (
	envPrefix = "DPM"
)

var (
	version string = "local"
	rootCmd        = command{
		command: &cobra.Command{
			Use:   "dpm",
			Short: "Development Project Manager",
			Long:  `A tool to manage development projects`,
			Run:   nil,
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				// Setup Viper
				err := viperSetup(cmd)
				if err != nil {
					return err
				}

				// Setup Logging
				setupLogging()

				return nil
			},
		},
		flags: func(cmd *cobra.Command) {
			cmd.PersistentFlags().String("config", "$HOME/.config/dpm/config.yml", "Location of configuration file")
			cmd.PersistentFlags().StringP("projects-home", "p", "$HOME/dev", "Projects location")
			cmd.PersistentFlags().BoolP("verbose", "v", false, "Enable more detailed logs")
		},
		subcommands: []command{
			{
				command: &cobra.Command{
					Use:   "clone",
					Short: "Clone a project",
					Long:  `Clone a git repo to the managed project directory`,
					RunE:  cloneCmd,
				},
				flags: func(cmd *cobra.Command) {
					cmd.PersistentFlags().BoolP("short", "s", false, "Output shortened to just project path")
					cmd.PersistentFlags().StringP("identity-file", "i", "$HOME/.ssh/id_rsa", "Path to the private key to use for git authentication")
				},
				subcommands: []command{},
			},
			{
				command: &cobra.Command{
					Use:   "list",
					Short: "List projects",
					Long:  `Lists all the known git repos`,
					RunE:  listCmd,
				},
				flags: func(cmd *cobra.Command) {
					cmd.PersistentFlags().BoolP("names", "n", false, "Only return project names")
				},
			},
		},
	}
)

func main() {
	// Create execution context
	ctx := context.Background()

	// Setup command structure
	err := registerCommands(nil, []command{rootCmd})
	if err != nil {
		slog.ErrorContext(
			ctx,
			"failed to process command structure",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	// Execute
	if err := fang.Execute(ctx, rootCmd.command); err != nil {
		slog.ErrorContext(
			ctx,
			"command failed",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
}

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
	if err := viper.ReadInConfig(); err == nil {
		return err
	}

	return nil
}

func setupLogging() {
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

// registerCommands recurses through the given command struct registering flags, and subcommands.
// `parentCommand` should be `nil` for the top level command.
func registerCommands(parentCommand *cobra.Command, commands []command) error {
	for _, cmd := range commands {
		// Register Flags
		cmd.flags(cmd.command)

		// Add command as subcommand
		if parentCommand != nil {
			parentCommand.AddCommand(cmd.command)
		}

		// Process subcommands
		err := registerCommands(cmd.command, cmd.subcommands)
		if err != nil {
			return err
		}
	}

	return nil
}
