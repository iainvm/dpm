package git

import (
	"regexp"
	"strings"
)

const (
	urlRegex = `^(([A-Za-z0-9]+@|http(|s)\:\/\/)|(http(|s)\:\/\/[A-Za-z0-9]+@))([A-Za-z0-9.]+(:\d+)?)(?::|\/)([\d\/\w.-]+?)(\.git){1}$`
	// ^(?:(ssh|git|http(?:s)?):\/\/)?
)

func GetUserFromURL(url string) string {
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
	site, _, _ := strings.Cut(groups[6], ":")

	// Get the rest of the subdirectories
	subdirs := strings.Split(groups[8], "/")

	// Add the elements to a slice
	elements := make([]string, 1+len(subdirs))
	elements[0] = site
	for i, dir := range subdirs {
		elements[i+1] = dir
	}

	return elements
}
