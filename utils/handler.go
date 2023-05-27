package utils

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  http.StatusNotFound,
		Message: "Page not found",
		Data:    nil,
	}
	SendJSONResponse(w, response, http.StatusNotFound)
}
