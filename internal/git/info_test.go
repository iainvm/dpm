package git

import (
	"testing"

	"github.com/spf13/viper"
)

func TestIsValidGitURL(t *testing.T) {
	testCases := map[string]bool{
		"git@github.com:iainvm/dev.git":     true,
		"https://github.com/iainvm/dev.git": true,
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
		"git@github.com:iainvm/dev.git":     "https://github.com/iainvm/dev.git",
		"https://github.com/iainvm/dev.git": "https://github.com/iainvm/dev.git",
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
		"git@github.com:iainvm/dev.git":     "/test_projects_home/github.com/iainvm/dev",
		"https://github.com/iainvm/dev.git": "/test_projects_home/github.com/iainvm/dev",
	}

	viper.Set("projects_home", "/test_projects_home")

	var result string
	var err error
	for url, expected := range testCases {
		result, err = GetProjectPath(url)
		if err != nil {
			t.Fail()
		}

		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	}
}
