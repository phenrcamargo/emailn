package middleware

import (
	"encoding/json"
	"net/http"
)

func SetDefaultDataTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := NewResponseRecorder(w)
		next.ServeHTTP(recorder, r)

		wrapped := map[string]interface{}{
			"statusCode": recorder.statusCode,
		}

		var data interface{}
		if json.Unmarshal(recorder.body.Bytes(), &data) == nil {
			wrapped["data"] = data
		} else {
			wrapped["data"] = recorder.body.String()
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(recorder.statusCode)
		json.NewEncoder(w).Encode(wrapped)
	})
}
