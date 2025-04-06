package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/iainvm/dpm/internal/config"
	"github.com/iainvm/dpm/internal/logger"
)

var flags struct {
	configPath string
	verbose    bool
}

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

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&flags.configPath, "config", "", "config file (default is $HOME/.config/dpm/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&flags.verbose, "verbose", "v", false, "log additional information")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Setup logger
	_ = logger.New(flags.verbose)

	// Setup config
	err := config.Load("DPM", flags.configPath)
	if err != nil {
		panic(err)
	}
}
