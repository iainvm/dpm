package git

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func OldClone(url string, location string, privateKeyLocation string) (*git.Repository, error) {
	_, err := os.Stat(privateKeyLocation)
	if err != nil {
		return nil, err
	}

	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyLocation, "")
	if err != nil {
		return nil, err
	}

	project, err := git.PlainClone(location, false, &git.CloneOptions{
		Auth:     publicKeys,
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to clone: %w", err)
	}

	return project, nil
}

func Clone(url string, target string, keyLocation string) (string, error) {
	// Check key exists

	// Load key from file

	// Clone project

	return "", nil
}
