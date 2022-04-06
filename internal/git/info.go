package git

import (
	"regexp"
	"strings"

	"github.com/spf13/viper"

	"github.com/iainvm/dev/internal/system"
)

// IsValidGitURL validates the given url against a RegEx to determine it's of an accepted format
func IsValidGitURL(url string) bool {
	var git_url_regex string = `^(([A-Za-z0-9]+@|http(|s)\:\/\/)|(http(|s)\:\/\/[A-Za-z0-9]+@))([A-Za-z0-9.]+(:\d+)?)(?::|\/)([\d\/\w.-]+?)(\.git){1}$`
	matched, _ := regexp.MatchString(git_url_regex, url)

	return matched
}

// TranslateToHTTP converts a given git url into a http version of it, for easier
func TranslateToHTTP(url string) string {
	// If already a http url return
	if strings.HasPrefix(url, "http") {
		return url
	}

	// Replace : after domain before adding the : used in the https://
	url = strings.Replace(url, ":", "/", 1)

	// Swap the git@ for use of https://
	if strings.HasPrefix(url, "git@") {
		url = strings.Replace(url, "git@", "https://", 1)
	}

	return url
}

// GetProjectPath deconstructs the url given and determines the directory path it for the project
// e.g. git@github.com:iainvm/dev.git     -> github.com/iainvm/dev
//      https://github.com/iainvm/dev.git -> github.com/iainvm/dev
func GetProjectPath(url string) (string, error) {
	var project_path string
	project_path = TranslateToHTTP(url)
	project_path = strings.Replace(project_path, "https://", "", 1)

	if strings.HasSuffix(project_path, ".git") {
		project_path = strings.Replace(project_path, ".git", "", 1)
	}

	var devDir string = viper.GetString("projects_home")

	var full_path string = devDir + "/" + project_path

	return system.AsAbsolutePath(full_path)
}
