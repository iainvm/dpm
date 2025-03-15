package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/internal/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dpm",
	Short: "A tool to manage development projects",
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

func verbosePrintf(location *os.File, message string, args ...interface{}) {
	verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
	if verbose {
		fmt.Fprintf(location, message, args...)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/dpm/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Log more details")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// Viper Config
	viper.SetDefault(config.KEY_PROJECTS_HOME, "~/dev")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		userHome, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in user's home directory with name ".dpm" (without extension).
		viper.AddConfigPath(fmt.Sprintf("%s/.config/dpm", userHome))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}
	viper.SetEnvPrefix(CONFIG_ENV_PREFIX)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		verbosePrintf(os.Stderr, "Using config file: %v\n", viper.ConfigFileUsed())
	}
}
