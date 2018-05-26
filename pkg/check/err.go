// Package check handles errors in defer statements
package check

import (
	"log"
)

// Err checks if deferredFunc throws an error and sets err if it hasn't already been set
func Err(deferredFunc func() error) {
	// func Err(deferredFunc func() error, err *error) {
	if err := deferredFunc(); err != nil {
		log.Printf("Error in deferred function: %v\n", err)
	}
}
