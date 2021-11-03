package host

import (
	"regexp"
	"strings"
)

const (
	githubURL        = "https://github.com"
	githubStorageURL = "https://codeload.github.com"
)

func URL(repo string) string {
	githubRegex := regexp.MustCompile(githubURL + `/(\S+)$`)
	if githubRegex.MatchString(repo) {
		return repo
	}

	return strings.Join([]string{githubURL, repo}, "/")
}
