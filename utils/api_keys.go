package utils

import (
	"net/http"
)

func CheckAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		expectedAPIKey := "CIMBNiaga" // Ganti dengan API key yang valid

		// Memeriksa keberadaan dan kevalidan API key
		if apiKey == "" || apiKey != expectedAPIKey {
			response := Response{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			}
			SendJSONResponse(w, response, http.StatusUnauthorized)
			return
		}

		// Lanjutkan ke handler selanjutnya jika API key valid
		next.ServeHTTP(w, r)
	})
}
