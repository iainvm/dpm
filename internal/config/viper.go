package config

import (
	"github.com/spf13/viper"

	"github.com/iainvm/dpm/internal/system"
)

func GetProjectsHome() (string, error) {
	path, err := getAsAbsolutePath(KEY_PROJECTS_HOME)
	return path, err
}

func GetPrivateKeyLocation() (string, error) {
	path, err := getAsAbsolutePath(KEY_PRIVATE_KEY_LOCATION)
	return path, err
}

func getAsAbsolutePath(key string) (string, error) {
	path := viper.GetString(key)
	path, err := system.AsAbsolutePath(path)

	return path, err
}
