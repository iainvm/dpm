package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/internal/config"
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
	Long: fmt.Sprintf(`Clone a git project into your projects home

	$ dpm clone git@github.com:iainvm/dpm.git
	Clones the project into $%s_%s/github.com/iainvm/dpm`, config.ENV_PREFIX, strings.ToUpper(config.PROJECTS_DIR)),
	Run: func(cmd *cobra.Command, args []string) {
		clone(args)
	},
}

func init() {
	cloneCmd.PersistentFlags().BoolP("short", "s", false, "Short output, will only return the absolute path to the project.")
	rootCmd.AddCommand(cloneCmd)
	viper.SetDefault(config.KEY_PRIVATE_KEY_LOCATION, "~/.ssh/id_rsa")
}

func clone(args []string) {
	ctx := context.Background()

	// Get args
	url := args[0]

	// Get Project Dir
	projectDir, err := config.ProjectsDir()
	cobra.CheckErr(err)

	// Get private key for git
	privateKeyLocation, err := config.GetPrivateKeyLocation()
	cobra.CheckErr(err)

	// Validate if project exists
	absoluteProjectPath := fmt.Sprintf("%s/%s", projectDir, git.URLAsPath(url))
	slog.DebugContext(
		ctx,
		"Path to clone to",
		slog.String("path", absoluteProjectPath),
	)

	// Check if project exists at that path
	projectExists, err := system.DoesFolderExist(absoluteProjectPath)
	cobra.CheckErr(err)
	if projectExists {
		fmt.Fprintf(os.Stdout, "Project already exists at: %s\n", absoluteProjectPath)
		os.Exit(0)
	}

	// Clone project if it doesn't exist
	_, err = git.Clone(url, absoluteProjectPath, privateKeyLocation)
	cobra.CheckErr(err)
}
