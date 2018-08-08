package validate

import "regexp"

const (
	projectRegexp = "^[\\w-]{5,30}$"
)

// ProjectRegexp returns the regexp to validate project strings
func ProjectRegexp() *regexp.Regexp {
	return regexp.MustCompile(projectRegexp)
}
