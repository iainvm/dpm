package git_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iainvm/dpm/internal/git"
)

func TestIsValidURL(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "HTTP URL",
			url:      "http://github.com/iainvm/dpm.git",
			expected: true,
		},
		{
			name:     "HTTPS URL",
			url:      "https://github.com/iainvm/dpm.git",
			expected: true,
		},
		{
			name:     "SSH URL",
			url:      "git@github.com:iainvm/dpm.git",
			expected: true,
		},
		{
			name:     "SSH URL with port",
			url:      "git@github.com:1234:iainvm/dpm.git",
			expected: true,
		},
		{
			name:     "Longer URL",
			url:      "user@gitlab.com:owner/a/b/c/project.git",
			expected: true,
		},
		{
			name:     "No user URL",
			url:      "gitlab.com:owner/a/b/c/project.git",
			expected: true,
		},
		{
			name:     "Slash not colon URL",
			url:      "git@gitlab.com/owner/a/b/c/project.git",
			expected: true,
		},
		{
			name:     "Wrong protocol URL",
			url:      "ftp://gitlab.com:owner/a/b/c/project.git",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := git.IsValidURL(tc.url)
			require.Equal(t, tc.expected, result)
		})
	}
}

func Ptr[T any](v T) *T {
	return &v
}

func TestSplitURL(t *testing.T) {
	testCases := []struct {
		name        string
		url         string
		directories git.URLInfo
	}{
		{
			name: "HTTP URL",
			url:  "http://github.com/iainvm/dpm.git",
			directories: git.URLInfo{
				Protocol: Ptr("http"),
				User:     nil,
				Host:     Ptr("github.com"),
				Port:     nil,
				Path:     Ptr("iainvm/dpm"),
			},
		},
		{
			name: "HTTPS URL",
			url:  "https://github.com/iainvm/dpm.git",
			directories: git.URLInfo{
				Protocol: Ptr("https"),
				User:     nil,
				Host:     Ptr("github.com"),
				Port:     nil,
				Path:     Ptr("iainvm/dpm"),
			},
		},
		{
			name: "SSH URL",
			url:  "git@github.com:iainvm/dpm.git",
			directories: git.URLInfo{
				Protocol: nil,
				User:     Ptr("git"),
				Host:     Ptr("github.com"),
				Port:     nil,
				Path:     Ptr("iainvm/dpm"),
			},
		},
		{
			name: "SSH URL with port",
			url:  "git@github.com:1234:iainvm/dpm.git",
			directories: git.URLInfo{
				Protocol: nil,
				User:     Ptr("git"),
				Host:     Ptr("github.com"),
				Port:     Ptr(1234),
				Path:     Ptr("iainvm/dpm"),
			},
		},
		{
			name: "Longer URL",
			url:  "user@gitlab.com:owner/a/b/c/project.git",
			directories: git.URLInfo{
				Protocol: nil,
				User:     Ptr("user"),
				Host:     Ptr("gitlab.com"),
				Port:     nil,
				Path:     Ptr("owner/a/b/c/project"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := git.GetInfo(tc.url)
			require.NoError(t, err)
			require.Equal(t, tc.directories, result)
		})
	}
}
