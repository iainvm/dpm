package cmd

import (
	"context"
	"fmt"
	"path"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"

	"github.com/iainvm/dpm/internal/config"
	"github.com/iainvm/dpm/internal/git"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects in projects home",
	Long:  `Searches the projects home directory for any git projects, returns the path of found git projects`,
	Run:   list,
}

func init() {
	listCmd.PersistentFlags().BoolP("name", "n", false, "Only return project names")
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	// Create context
	ctx := context.Background()
	log := slog.Default()

	// Get Flags
	nameOnly, err := cmd.PersistentFlags().GetBool("name")
	cobra.CheckErr(err)
	log.Info("Info log line")
	log.Debug("Debug log line")
	log.DebugContext(
		ctx,
		"Flag values",
		// slog.Group("flags",
		// 	slog.Bool("name", nameOnly),
		// ),
	)
	log.InfoContext(ctx, "Info log line with ctx")

	// Get projects directory
	projectsDir, err := config.ProjectsDir()
	cobra.CheckErr(err)

	// Get all git projects in path
	projectPaths, err := git.SearchPathForGit(projectsDir)
	cobra.CheckErr(err)

	// Display result
	for _, projectPath := range projectPaths {
		if nameOnly {
			fmt.Println(path.Base(projectPath))
		} else {
			fmt.Println(projectPath)
		}
	}
}
