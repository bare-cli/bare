package host

import (
	"regexp"
	"strings"
)

const (
	githubURL        = "https://github.com"
	githubStorageURL = "https://codeload.github.com"
)

// func ZipURL(repo string) string {
// 	var version = "main"

// 	repo = strings.TrimSuffix(strings.TrimPrefix(repo, "/"), "/")

// 	zipRegex := regexp.MustCompile(`zip/(\S+)$`)
// 	if zipRegex.MatchString(repo) {
// 		return repo
// 	}

// }

func URL(repo string) string {
	githubRegex := regexp.MustCompile(githubURL + `/(\S+)$`)
	if githubRegex.MatchString(repo) {
		return repo
	}

	return strings.Join([]string{githubURL, repo}, "/")
}
