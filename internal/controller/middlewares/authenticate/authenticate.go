package authenticate

import (
	"log/slog"
	"net/http"
)

func WithApiKey() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-KEY")
			// TODO: Create Service that generates Key and put them in DB
			if apiKey != "123" {
				slog.Info("API Key is not valid")
				authFailed(w)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func authFailed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}
