package git_test

import (
	"testing"

	"github.com/iainvm/dpm/internal/git"
	"github.com/stretchr/testify/require"
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
			expected: false,
		},
		{
			name:     "Wrong protocol URL",
			url:      "ftp://gitlab.com:owner/a/b/c/project.git",
			expected: false,
		},
		{
			name:     "Slash not colon URL",
			url:      "git@gitlab.com/owner/a/b/c/project.git",
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

func TestSplitURL(t *testing.T) {
	testCases := []struct {
		name        string
		url         string
		directories []string
	}{
		{
			name:        "HTTP URL",
			url:         "http://github.com/iainvm/dpm.git",
			directories: []string{"github.com", "iainvm", "dpm"},
		},
		{
			name:        "HTTPS URL",
			url:         "https://github.com/iainvm/dpm.git",
			directories: []string{"github.com", "iainvm", "dpm"},
		},
		{
			name:        "SSH URL",
			url:         "git@github.com:iainvm/dpm.git",
			directories: []string{"github.com", "iainvm", "dpm"},
		},
		{
			name:        "SSH URL with port",
			url:         "git@github.com:1234:iainvm/dpm.git",
			directories: []string{"github.com", "iainvm", "dpm"},
		},
		{
			name:        "Longer URL",
			url:         "user@gitlab.com:owner/a/b/c/project.git",
			directories: []string{"gitlab.com", "owner", "a", "b", "c", "project"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := git.SplitURL(tc.url)
			require.Equal(t, tc.directories, result)
		})
	}
}
