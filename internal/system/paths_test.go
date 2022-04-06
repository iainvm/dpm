package system

import (
	"testing"
)

func mockHomeDir() (string, error) {
	return "/home/dir", nil
}

func mockWorkDir() (string, error) {
	return "/work/dir", nil
}

func TestCreatePathFromWorkingDirectory(t *testing.T) {
	// mock methods
	homeDir = mockHomeDir
	workDir = mockWorkDir

	testCases := map[string]string{
		"~/path/to/file":         "/home/dir/path/to/file",
		"path/to/file":           "/work/dir/path/to/file",
		"/absolute/path/to/file": "/absolute/path/to/file",
	}

	for path, expected := range testCases {
		result := AsAbsolutePath(path)

		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	}
}
