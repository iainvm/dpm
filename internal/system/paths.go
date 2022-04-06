package system

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	homeDir = os.UserHomeDir
	workDir = os.Getwd
)

func AsAbsolutePath(path string) string {
	switch {
	case strings.HasPrefix(path, "~"):
		return replaceTildeWithHome(path)
	case !strings.HasPrefix(path, "/"):
		return createPathFromWorkingDirectory(path)
	default:
		return path
	}
}

func createPathFromWorkingDirectory(path string) string {
	workingDirectory, err := workDir()
	cobra.CheckErr(err)

	return workingDirectory + "/" + path
}

func replaceTildeWithHome(path string) string {
	home, err := homeDir()
	cobra.CheckErr(err)

	path = strings.Replace(path, "~", home, 1)

	return path
}
