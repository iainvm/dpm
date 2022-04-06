package git

import (
	"log"
)

func Clone(url string) error {
	projectPath := GetProjectPath(url)

	log.Printf("%s", projectPath)

	return nil
}
