package dpm

import (
	"context"
	"fmt"
	"log/slog"
	"path"

	"github.com/iainvm/dpm/internal/git"
)

// Clone takes a url, and authFile to use git to clone a repository
func Clone(ctx context.Context, devDir string, url string) (string, error) {
	// Check URL is valid
	if !git.IsValidURL(url) {
		return "", fmt.Errorf("received invalid git URL: %s", url)
	}

	// Calculate directory to clone to
	urlInfo, err := git.GetInfo(url)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"failed to get repo info",
			slog.Any("error", err),
		)
		return "", fmt.Errorf("failed to get repo info: %w", err)
	}

	directory := path.Join(devDir, *urlInfo.Host, *urlInfo.Path)

	slog.DebugContext(
		ctx,
		"DPM Clone called",
		slog.String("url", url),
		slog.String("dev_dir", devDir),
		slog.String("directory", directory),
	)

	// Clone the repo to the directory
	_, err = git.Clone(ctx, url, directory)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"failed to clone git repo",
			slog.Any("error", err),
		)
		return "", fmt.Errorf("failed to clone git repo to '%s': %w", directory, err)
	}

	return directory, nil
}
