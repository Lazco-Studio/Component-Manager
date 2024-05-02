package module

import (
	"regexp"
)

func GithubNamingRule(s string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)
	return regex.MatchString(s)
}
