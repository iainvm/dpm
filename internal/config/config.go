package config

import (
	"github.com/iainvm/dpm/internal/system"
	"github.com/spf13/viper"
)

// List of configuration keys

const (
	ENV_PREFIX               string = "DPM"
	KEY_PROJECTS_HOME        string = "projects_home"
	KEY_PRIVATE_KEY_LOCATION string = "private_key_location"
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
