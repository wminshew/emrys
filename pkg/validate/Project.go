package validate

import "regexp"

const (
	projectRe = "^[\\w-]{5,30}$"
)

// ProjectRe returns the regexp to validate project strings
func ProjectRe() *regexp.Regexp {
	return regexp.MustCompile(projectRe)
}
