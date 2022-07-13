package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/internal/git"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects in projects home",
	Long:  `Searches the projects home directory for any git projects, returns the path of found git projects`,
	Run: func(cmd *cobra.Command, args []string) {
		list(args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("name", "n", false, "Only return project names")
}

func list(args []string) {
	projects_home := viper.GetString(CONFIG_KEY_PROJECTS_HOME)
	nameOnly, err := rootCmd.PersistentFlags().GetBool("name")
	cobra.CheckErr(err)

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
