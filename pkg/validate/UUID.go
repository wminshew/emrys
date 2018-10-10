package validate

import (
	"fmt"
	"regexp"
)

const (
	uuidV4Regexp = "[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}"
)

// UUIDRegexp returns the regexp to validate uuid strings in the client
func UUIDRegexp() *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("^%s$", uuidV4Regexp))
}

// UUIDRegexpMux returns the regexp to validate uuid paths with gorilla mux
func UUIDRegexpMux() *regexp.Regexp {
	return regexp.MustCompile(uuidV4Regexp)
}
