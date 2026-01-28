package git

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hashicorp/go-getter"
)

// Clone will clone the git repository to the given directory
func Clone(ctx context.Context, url string, directory string) error {
	slog.DebugContext(
		ctx,
		"Cloning git repository",
		slog.String("url", url),
		slog.String("target", directory),
	)

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
