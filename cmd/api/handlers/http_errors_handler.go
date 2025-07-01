package handlers

import (
	"emailn/internal/shared/exception"
	"encoding/json"
	"net/http"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func ErrorHandler(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)

		if err != nil {
			if e, ok := err.(exception.HttpNotFoundException); ok {
				http.Error(w, e.Message, http.StatusNotFound)
				return
			}

			if e, ok := err.(exception.HttpBadRequestException); ok {
				http.Error(w, e.Message, http.StatusBadRequest)
				return
			}

			if e, ok := err.(exception.HttpInternalServerException); ok {
				http.Error(w, e.Message, http.StatusInternalServerError)
				return
			}

			http.Error(w, "An internal server error ocurred", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(status)
		if obj != nil {
			json.NewEncoder(w).Encode(obj)
		}
	})
}
