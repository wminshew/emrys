package validate

import (
	"regexp"
)

const (
	uuidV4Regexp = "[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}"
)

// UUIDRegexpMux returns the regexp to validate uuid paths with gorilla mux
func UUIDRegexpMux() *regexp.Regexp {
	return regexp.MustCompile(uuidV4Regexp)
}
