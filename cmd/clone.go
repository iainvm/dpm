package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/internal/git"
	"github.com/iainvm/dpm/internal/system"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a git project into your DPM_ROOT",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("only accepts 1 argument")
		}
		if !git.IsValidGitURL(args[0]) {
			return errors.New("not a valid git url")
		}
		return nil
	},
	Long: `Will clone a git project into your projects home

	e.g. git@github.com:iainvm/dpm.git
	will clone to
	$DPM_PROJECTS_HOME/github.com/iainvm/dpm`,
	Run: func(cmd *cobra.Command, args []string) {
		clone(args)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func clone(args []string) {
	// Process args
	url := args[0]
	projects_home := viper.GetString(PROJECTS_HOME)

	// Get additional info
	var projectPath string = git.GetProjectPath(url)
	absoluteProjectPath, err := system.AsAbsolutePath(projects_home + "/" + projectPath)
	cobra.CheckErr(err)
	verbosePrintf(os.Stdout, "Project path: %s\n", absoluteProjectPath)

	// Validate if project exists
	projectExists, err := system.DoesFolderExist(absoluteProjectPath)
	cobra.CheckErr(err)
	if projectExists {
		fmt.Fprintf(os.Stdout, "Project already exists at: %s", absoluteProjectPath)
		os.Exit(0)
	}

	// Clone project if it doesn't exist
	_, err = git.Clone(url, absoluteProjectPath)
	cobra.CheckErr(err)
}
