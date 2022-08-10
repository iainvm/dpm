package git

import (
	"regexp"
	"strings"

	"github.com/iainvm/dpm/internal/system"
)

// IsValidGitURL validates the given url against a RegEx to determine it's of an accepted format
func IsValidGitURL(url string) bool {
	var git_url_regex string = `^(([A-Za-z0-9]+@|http(|s)\:\/\/)|(http(|s)\:\/\/[A-Za-z0-9]+@))([A-Za-z0-9.]+(:\d+)?)(?::|\/)([\d\/\w.-]+?)(\.git){1}$`
	matched, _ := regexp.MatchString(git_url_regex, url)

	return matched
}

func IsGitProjectPath(path string) (bool, error) {
	return system.DoesFolderExist(path + "/.git")
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
// e.g. git@github.com:iainvm/dpm.git     -> github.com/iainvm/dpm
//      https://github.com/iainvm/dpm.git -> github.com/iainvm/dpm
func GetProjectPath(url string) string {
	var project_path string
	project_path = TranslateToHTTP(url)
	project_path = strings.Replace(project_path, "https://", "", 1)

	if strings.HasSuffix(project_path, ".git") {
		project_path = strings.Replace(project_path, ".git", "", 1)
	}

	return project_path
}

func GetProjectPathsInPath(path string) ([]string, error) {
	var gitProjects []string

	folders, err := system.GetDirectoriesInPath(path)
	if err != nil {
		return nil, err
	}

	for _, folder := range folders {
		if strings.HasSuffix(folder, ".git") {
			gitProjects = append(gitProjects, strings.TrimSuffix(folder, "/.git"))
		}
	}
	return gitProjects, nil
}
