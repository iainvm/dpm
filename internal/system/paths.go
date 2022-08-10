package system

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	homeDir = os.UserHomeDir
	workDir = os.Getwd

	homeDirShorthand    = "~"
	pathSeparator       = string(os.PathSeparator)
	currentDirShorthand = fmt.Sprintf(".%s", string(os.PathSeparator))
	parentDirShorthand  = fmt.Sprintf("..%s", string(os.PathSeparator))
)

func AsAbsolutePath(path string) (string, error) {
	switch {
	case strings.HasPrefix(path, homeDirShorthand):
		return replaceTildeWithHome(path)
	case !strings.HasPrefix(path, pathSeparator):
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

	for strings.HasPrefix(path, parentDirShorthand) {
		workingDirectory = filepath.Dir(workingDirectory)
		path = strings.TrimPrefix(path, parentDirShorthand)
	}

	path = strings.TrimPrefix(path, currentDirShorthand)
	return workingDirectory + pathSeparator + path, nil
}

func replaceTildeWithHome(path string) (string, error) {
	home, err := homeDir()
	if err != nil {
		return "", err
	}

	path = strings.Replace(path, homeDirShorthand, home, 1)

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
