package job

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes the value v to the http response stream as json with standard json encoding
// Source: https://github.com/moby/moby/blob/3a633a712c8bbb863fe7e57ec132dd87a9c4eff7/api/server/httputils/httputils_write_json.go
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	// enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
