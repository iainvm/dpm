/*
Copyright © 2022 Iain Majer iainvm@outlook.com

*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/iainvm/dev/internal/git"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a git project into your DEV_ROOT",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("only accepts 1 argument")
		}
		if !git.IsValidGitURL(args[0]) {
			return errors.New("not a valid git url")
		}
		return nil
	},
	Long: `Will clone a git project into your DEV_ROOT

	e.g. git@github.com:iainvm/dev.git
	will clone to
	$DEV_ROOT/github.com/iainvm/dev`,
	Run: func(cmd *cobra.Command, args []string) {
		exec(args)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func exec(args []string) {
}
