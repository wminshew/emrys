package validate

import (
	"regexp"
)

const (
	relPathRegexp     = "[0-9A-Za-z_ .\\\\/-]"
	relPathAntiRegexp = "\\.\\."
)

// RelPathRegexp returns a regexp to validate relative paths within data directories
func RelPathRegexp() *regexp.Regexp {
	return regexp.MustCompile(relPathRegexp)
}

// RelPathAntiRegexp returns a regexp to validate relative paths within data directories
func RelPathAntiRegexp() *regexp.Regexp {
	return regexp.MustCompile(relPathAntiRegexp)
}
