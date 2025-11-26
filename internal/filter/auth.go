package filter

import (
	"food-app/internal/utils"
	"net/http"
)

func ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api_key")
		if apiKey == "" {
			apiKey = r.Header.Get("Api-Key")
		}
		if apiKey == "" {
			apiKey = r.Header.Get("API-Key")
		}

		if apiKey != "apitest" {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid or missing API key")
			return
		}
		next.ServeHTTP(w, r)
	})
}
