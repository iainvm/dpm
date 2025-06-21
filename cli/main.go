package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

type command struct {
	command     *cobra.Command
	flags       func(parentCmd *cobra.Command, cmd *cobra.Command)
	subcommands []command
}

var (
	cfgFile string
	rootCmd = command{
		command: &cobra.Command{
			Use:   "dpm",
			Short: "Development Project Manager",
			Long:  `A tool to manage development projects`,
			Run:   nil,
		},
		flags: func(parentCmd *cobra.Command, cmd *cobra.Command) {
			cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/dpm/config.yaml)")
			cmd.PersistentFlags().BoolP("verbose", "v", false, "Log more details")
		},
		subcommands: []command{
			{
				command: &cobra.Command{
					Use:   "list",
					Short: "List projects",
					Long:  `Lists all the known git repos`,
					Run: func(cmd *cobra.Command, args []string) {
						slog.Info("list command executed")
						// parse args
						// call dpm.list
					},
				},
				flags: func(parentCmd *cobra.Command, cmd *cobra.Command) {
					cmd.PersistentFlags().BoolP("name", "n", false, "Only return project names")
				},
			},
		},
	}
)

func setupCommands(parentCommand *cobra.Command, commands []command) error {
	for _, cmd := range commands {
		// Register Flags
		cmd.flags(parentCommand, cmd.command)

		// Add command as subcommand
		if parentCommand != nil {
			parentCommand.AddCommand(cmd.command)
		}

		// Process subcommands
		err := setupCommands(cmd.command, cmd.subcommands)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := setupCommands(nil, []command{rootCmd})
	if err != nil {
		slog.Error(
			"failed to process command structure",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	// Execute
	err = rootCmd.command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
