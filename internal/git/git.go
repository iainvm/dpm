package git

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

// Clone will git clone the git repository to the given directory, using the keyLocation and keyPassword for authentication
func Clone(ctx context.Context, url string, directory string, silent bool) (*git.Repository, error) {
	slog.DebugContext(
		ctx,
		"Cloning git repository",
		slog.String("url", url),
		slog.String("target", directory),
		slog.Bool("log", silent),
	)

	// Use SSH Agent
	user := GetUserFromURL(url)
	authMethod, err := ssh.NewSSHAgentAuth(user)
	if err != nil {
		return nil, fmt.Errorf("failed to access SSH Agent: %w", err)
	}

	// Determine if progress should be logged
	var progress io.Writer = nil
	if !silent {
		progress = os.Stdout
	}

	// Clone project
	project, err := git.PlainClone(directory, &git.CloneOptions{
		Auth:     authMethod,
		URL:      url,
		Progress: progress,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to clone: %w", err)
	}

	return project, nil
}
