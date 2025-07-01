package middleware

import (
	"net/http"
	"strings"
)

func RequiredJSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if !strings.HasPrefix(contentType, "application/json") {
			http.Error(w, "Content-Type must be JSON", http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	})
}
