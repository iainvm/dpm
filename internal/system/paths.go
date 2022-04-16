package system

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	homeDir = os.UserHomeDir
	workDir = os.Getwd
)

func AsAbsolutePath(path string) (string, error) {
	switch {
	case strings.HasPrefix(path, "~"):
		return replaceTildeWithHome(path)
	case !strings.HasPrefix(path, "/"):
		return createPathFromWorkingDirectory(path)
	default:
		return path, nil
	}
}

func createPathFromWorkingDirectory(path string) (string, error) {
	workingDirectory, err := workDir()
	if err != nil {
		return "", err
	}

	return workingDirectory + "/" + path, nil
}

func replaceTildeWithHome(path string) (string, error) {
	home, err := homeDir()
	if err != nil {
		return "", err
	}

	path = strings.Replace(path, "~", home, 1)

	return path, nil
}

func DoesFolderExist(path string) (bool, error) {
	_, err := os.Stat(path)
	switch {
	case err == nil:
		return true, nil
	case os.IsNotExist(err):
		return false, nil
	default:
		return false, err
	}
}

func GetDirectoriesInPath(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
