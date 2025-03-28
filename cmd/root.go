package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/internal/config"
)

var configPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dpm",
	Short: "A tool to manage development projects",
	Long:  `A tool to manage development projects`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func verbosePrintf(location *os.File, message string, args ...interface{}) {
	verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
	if verbose {
		fmt.Fprintf(location, message, args...)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "config file (default is $HOME/.config/dpm/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Log more details")

	// Viper Config
	viper.SetDefault(config.KEY_PROJECTS_HOME, "~/dev")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	err := config.Load("DPM", configPath)
	if err != nil {
		panic(err)
	}
}
