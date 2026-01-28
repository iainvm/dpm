package main

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"

	"github.com/iainvm/dpm/internal/logger"
)

type command struct {
	command     *cobra.Command
	flags       func(cmd *cobra.Command)
	subcommands []command
}

var (
	version   string = "local"
	buildDate string = "unknown"

	output io.Writer = os.Stdout

	rootCmd = command{
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
				logger.NewCLILogger(
					slog.Group(
						"build",
						slog.String("version", version),
						slog.String("date", buildDate),
					),
				)

				return nil
			},
		},
		flags: func(cmd *cobra.Command) {
			cmd.PersistentFlags().String("config", "$HOME/.config/dpm/config.yml", "Location of configuration file")
			cmd.PersistentFlags().StringP("dev-directory", "d", "$HOME/dev", "Where all the dev projects go")
			cmd.PersistentFlags().BoolP("verbose", "v", false, "Enable more detailed logs")
		},
		subcommands: []command{
			{
				command: &cobra.Command{
					Use:     "clone",
					Short:   "Clone a project",
					Long:    `Clone a git repo to the managed project directory`,
					Example: "dpm clone git@github.com:iainvm/dpm.git",
					Args:    cobra.MatchAll(cobra.ExactArgs(1)),
					RunE:    cloneCmd,
				},
				flags:       func(cmd *cobra.Command) {},
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

	// Set version
	rootCmd.command.SetVersionTemplate(version)

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
		os.Exit(1)
	}
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
