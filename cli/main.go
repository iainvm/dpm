package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

type command struct {
	command     *cobra.Command
	flags       func(cmd *cobra.Command)
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
		flags: func(cmd *cobra.Command) {
			cmd.PersistentFlags().StringVar(&cfgFile, "config", "~/.config/dpm/config.yaml", "Location of configuration file")
			cmd.PersistentFlags().BoolP("verbose", "v", false, "Enable more detailed logs")
		},
		subcommands: []command{
			{
				command: &cobra.Command{
					Use:   "clone",
					Short: "Clone a project",
					Long:  `Clone a git repo to the managed project directory`,
					Run: func(cmd *cobra.Command, args []string) {
						slog.Info("clone command executed")
						// parse args
						// call dpm.clone
					},
				},
				flags: func(cmd *cobra.Command) {
					cmd.PersistentFlags().BoolP("short", "s", false, "Output shortened to just project path")
					cmd.PersistentFlags().String("ssh-key-path", "~/.ssh/id_rsa", "Path to the private key to use for git authentication")
				},
			},
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
				flags: func(cmd *cobra.Command) {
					cmd.PersistentFlags().BoolP("name", "n", false, "Only return project names")
				},
			},
		},
	}
)

func setupCommands(parentCommand *cobra.Command, commands []command) error {
	for _, cmd := range commands {
		// Register Flags
		cmd.flags(cmd.command)

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
	if err := fang.Execute(context.Background(), rootCmd.command); err != nil {
		os.Exit(1)
	}
}
