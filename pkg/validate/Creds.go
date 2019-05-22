package validate

import (
	"regexp"
	"unicode"
)

const (
	emailRegexp = "@"
)

// EmailRegexp returns the regexp to validate email strings on the server
func EmailRegexp() *regexp.Regexp {
	return regexp.MustCompile(emailRegexp)
}

// Password returns whether the password is valid
func Password(s string) bool {
	var (
		hasMinLen  = false
		hasMaxLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	if len(s) <= 30 {
		hasMaxLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasMaxLen && hasUpper && hasLower && hasNumber && hasSpecial
}
