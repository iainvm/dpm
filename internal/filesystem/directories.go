package filesystem

import (
	"io"
	"os"
)

func IsExists(directory string) bool {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsEmpty(directory string) (bool, error) {
	if !IsExists(directory) {
		return true, nil
	}

	f, err := os.Open(directory)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
