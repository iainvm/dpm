package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

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
	projects_home, err := config.GetProjectsHome()
	cobra.CheckErr(err)

	nameOnly, err := cmd.PersistentFlags().GetBool("name")
	cobra.CheckErr(err)
	verbosePrintf(os.Stdout, "Name flag: %v\n", nameOnly)

	projectPaths, err := git.GetProjectPathsInPath(projects_home)
	cobra.CheckErr(err)
	for _, projectPath := range projectPaths {
		if nameOnly {
			index := strings.LastIndex(projectPath, "/")
			fmt.Println(projectPath[index+1:])
		} else {
			fmt.Println(projectPath)
		}
	}
}
