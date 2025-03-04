package handlers

import (
	"encoding/json"
	"net/http"
)

// reportError helper to return error as response.
func reportError(w http.ResponseWriter, code int, desc string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		ErrorDescription: desc,
	})
}

// ErrorResponse provides structured information related to errors
// returned by the server.
type ErrorResponse struct {
	// Human-readable text providing
	// additional information, used to assist the client developer in
	// understanding the error that occurred.
	ErrorDescription string `json:"error_description,omitempty"`
}
