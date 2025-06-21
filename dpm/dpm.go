package dpm

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dpm",
	Short: "Development Project Manager",
	Long:  `A tool to manage development projects`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

func init() {
	// initConfig()
	// cobra.OnInitialize(initConfig)

	// Flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/dpm/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Log more details")

	// Config
	// viper.SetDefault(config.KEY_PROJECTS_HOME, "~/dev")
}
