package dpm

import (
	"context"
	"os"
	"path"
	"path/filepath"
)

const gitDirName = ".git"

type ListOptions struct {
	Names bool
}

// List will print out a list of paths or names of all the git repos it can find inside the dev directory
func List(ctx context.Context, devDir string, options *ListOptions) ([]string, error) {
	list := []string{}

	err := filepath.Walk(devDir, func(filePath string, fi os.FileInfo, err error) error {
		// If it's a `.git` dir
		if fi.Name() == gitDirName {

			// Get the parent dirs path
			elem := path.Dir(filePath)
			// If we want only the name, prune it
			if options.Names {
				elem = path.Base(elem)
			}

			// Add entry to the list
			list = append(list, elem)

			// Return and don't check inside the `.git` dir
			return filepath.SkipDir
		}

		return nil
	})

	return list, err
}
