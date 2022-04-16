/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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
	Long: `List all projects that have been cloned into the projects home folder	`,
	Run: func(cmd *cobra.Command, args []string) {
		list(args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.PersistentFlags().BoolP("name", "n", false, "Log more details")
}

func list(args []string) {
	var projects_home string = viper.GetString(PROJECTS_HOME)
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
