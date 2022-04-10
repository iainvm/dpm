package git

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func Clone(url string, location string) (*git.Repository, error) {

	project, err := git.PlainClone(location, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, err
	}

	return project, nil
}
