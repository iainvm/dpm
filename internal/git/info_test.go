package git

import (
	"testing"
)

func TestIsValidGitURL(t *testing.T) {
	testCases := map[string]bool{
		"git@github.com:iainvm/dpm.git":     true,
		"https://github.com/iainvm/dpm.git": true,
		"google.com":                        false,
	}

	for url, expected := range testCases {
		valid := IsValidGitURL(url)

		if valid != expected {
			t.Errorf("Expected %t, but got %t", expected, valid)
		}
	}
}

func TestTranslateToHTTP(t *testing.T) {
	testCases := map[string]string{
		"git@github.com:iainvm/dpm.git":     "https://github.com/iainvm/dpm.git",
		"https://github.com/iainvm/dpm.git": "https://github.com/iainvm/dpm.git",
	}

	var result string
	for url, expected := range testCases {
		result = TranslateToHTTP(url)

		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	}
}

func TestGetProjectPath(t *testing.T) {
	testCases := map[string]string{
		"git@github.com:iainvm/dpm.git":     "github.com/iainvm/dpm",
		"https://github.com/iainvm/dpm.git": "github.com/iainvm/dpm",
	}

	var result string
	for url, expected := range testCases {
		result = GetProjectPath(url)

		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	}
}
