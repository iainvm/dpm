package git

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hashicorp/go-getter"

	"github.com/iainvm/dpm/internal/filesystem"
)

// Clone will clone the git repository to the given directory
func Clone(ctx context.Context, url string, directory string) error {
	slog.DebugContext(
		ctx,
		"Cloning git repository",
		slog.String("url", url),
		slog.String("target", directory),
	)

	exists := filesystem.IsExists(directory)
	if exists {
		return fmt.Errorf("cannot clone, directory already exists: %s", directory)
	}

	// Create client to clone given project
	client := getter.Client{
		Ctx:  ctx,
		Src:  fmt.Sprintf("git::%s", url),
		Dst:  directory,
		Mode: getter.ClientModeDir,
	}

	// Clone repo and return any errors
	err := client.Get()
	if err != nil {
		return fmt.Errorf("failed to clone project: %w", err)
	}

	return nil
}
