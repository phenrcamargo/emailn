package api

import (
	"emailn/cmd/api/middleware"
	campaignRoutes "emailn/internal/campaign/presentation/routes"

	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequiredJSONContentType)
	r.Use(middleware.SetDefaultDataTypeMiddleware)

	r.Route("/campaign", func(r chi.Router) {
		r.Mount("/", campaignRoutes.Routes())
	})

	return r
}
