// Package check handles errors in defer statements
package check

import (
	"log"
)

// Err logs if deferredFunc throws an error; primarily used for deferred funcs
func Err(errorFunc func() error) {
	if err := errorFunc(); err != nil {
		log.Printf("check.Err: %v\n", err)
	}
}
