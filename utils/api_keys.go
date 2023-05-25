package utils

import (
	"encoding/json"
	"net/http"
)

const (
	APIKeyHeader = "X-API-Key"
	ValidAPIKey  = "CIMBNiaga"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

var (
	ErrInvalidAPIKey = ErrorResponse{Message: "Invalid API Key"}
)

func ValidateAPIKey(r *http.Request) error {
	apiKey := r.Header.Get(APIKeyHeader)
	if apiKey != ValidAPIKey {
		return ErrInvalidAPIKey
	}
	return nil
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
