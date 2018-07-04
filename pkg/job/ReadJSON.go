package job

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// ReadJSON reads the reader r as standard json and prints the stream to log
func ReadJSON(r io.Reader) error {
	var stream map[string]interface{}

	dec := json.NewDecoder(r)
	for dec.More() {
		if err := dec.Decode(&stream); err != nil {
			return fmt.Errorf("Error decoding json stream: %v", err)
		}
		for k, v := range stream {
			if k == "error" {
				return fmt.Errorf("%v", v)
			} else if k == "stream" {
				log.Printf("%v", v)
			} else {
				log.Printf("%v: %v\n", k, v)
			}
		}
	}
	return nil
}
