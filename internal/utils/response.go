package utils

import (
	"encoding/json"
	"food-app/internal/models"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, models.ApiResponse{
		Code:    code,
		Type:    "error",
		Message: message,
	})
}
