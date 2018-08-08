package validate

import (
	"fmt"
	"regexp"
)

const (
	projectRegexpMux = "[\\w-]{5,30}"
)

// ProjectRegexp returns the regexp to validate project strings in the client
func ProjectRegexp() *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("^%s$", projectRegexpMux))
}

// ProjectRegexpMux returns the regexp to validate project paths with gorilla mux
func ProjectRegexpMux() *regexp.Regexp {
	return regexp.MustCompile(projectRegexpMux)
}
