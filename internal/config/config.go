package config

import (
	"github.com/iainvm/dpm/internal/system"
	"github.com/spf13/viper"
)

// List of configuration keys

const (
	ENV_PREFIX               string = "DPM"
	PROJECTS_DIR             string = "projects_dir"
	KEY_PRIVATE_KEY_LOCATION string = "private_key_location"
)

func ProjectsDir() (string, error) {
	path, err := getAsAbsolutePath(PROJECTS_DIR)
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
