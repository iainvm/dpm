package git

import (
	"regexp"
	"strings"
)

const (
	// Capture Groups: protocol, user, host, port, path
	urlRegex = `^(?:(ssh|git|https?|git):\/\/)?(?:([\w\-]+)@)?([\w\.\-]+)(?::(\d+))?[:/]((?:~?[\w\-]+\/)?[\w\-./]+(?:\.git)?)\/?$`
)

func GetUserFromURL(url string) string {
	r := regexp.MustCompile(urlRegex)

	// Find matching groups
	groups := r.FindStringSubmatch(url)
	return groups[2]
}

// IsValidURL validates the given url against a RegEx to determine it's of an accepted format
func IsValidURL(url string) bool {
	r := regexp.MustCompile(urlRegex)
	matched := r.MatchString(url)
	return matched
}

// SplitURL will parse the git repository URL and break it down into URL, groups, and project
//
//	"git@github.com:iainvm/dpm.git" => ["github.com", "iainvm", "dpm"]
//	"git@gitlab.com:company/product/project.git" => ["gitlab.com", "company", "product", "project"]
func SplitURL(url string) []string {
	// Initialise regex
	r := regexp.MustCompile(urlRegex)

	// Find matching groups
	groups := r.FindStringSubmatch(url)

	// Get the website from the groups
	site := groups[3]

	// Get the rest of the subdirectories
	path, _ := strings.CutSuffix(groups[5], ".git")
	subdirs := strings.Split(path, "/")

	// Add the elements to a slice
	elements := make([]string, 1+len(subdirs))
	elements[0] = site
	for i, dir := range subdirs {
		elements[i+1] = dir
	}

	return elements
}
