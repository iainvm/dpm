package git

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	// Capture Groups: protocol, user, host, port, path
	urlRegex = `^(?:(ssh|git|https?|git):\/\/)?(?:([\w\-]+)@)?([\w\.\-]+)(?::(\d+))?[:/]((?:~?[\w\-]+\/)?[\w\-./]+(?:\.git)?)\/?$`
)

type URLInfo struct {
	Protocol *string
	User     *string
	Host     *string
	Port     *int
	Path     *string
}

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

// GetInfo will parse the git repository URL and determine the protocol, user, host, port, and path
func GetInfo(url string) (URLInfo, error) {
	// Initialise regex
	r := regexp.MustCompile(urlRegex)

	// Find matching groups
	groups := r.FindStringSubmatch(url)

	// Parse protocol
	var protocol *string
	if groups[1] != "" {
		protocol = &groups[1]
	}

	// Parse user
	var user *string
	if groups[2] != "" {
		user = &groups[2]
	}

	// Parse host
	var host *string
	if groups[3] != "" {
		host = &groups[3]
	}

	// Parse port
	var port *int
	if groups[4] != "" {
		data, err := strconv.Atoi(groups[4])
		if err != nil {
			return URLInfo{}, fmt.Errorf("failed to parse repo path: %w", err)
		}
		port = &data
	}

	// Parse Path
	var path *string
	if groups[5] != "" {
		data, _ := strings.CutSuffix(groups[5], ".git")
		path = &data
	}

	split := URLInfo{
		Protocol: protocol,
		User:     user,
		Host:     host,
		Port:     port,
		Path:     path,
	}

	return split, nil
}
