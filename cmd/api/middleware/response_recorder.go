package middleware

import (
	"bytes"
	"net/http"
)

type ResponseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           &bytes.Buffer{},
	}
}

func (w *ResponseRecorder) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *ResponseRecorder) Write(data []byte) (int, error) {
	return w.body.Write(data)
}
