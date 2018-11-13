package validate

import (
	"regexp"
)

const (
	emailRegexp    = "@"
	passwordRegexp = "^[A-Za-z0-9]{8, 30}$"
)

// EmailRegexp returns the regexp to validate email strings on the server
func EmailRegexp() *regexp.Regexp {
	return regexp.MustCompile(emailRegexp)
}

// PasswordRegexp returns the regexp to validate password strings on the server
func PasswordRegexp() *regexp.Regexp {
	return regexp.MustCompile(passwordRegexp)
}
