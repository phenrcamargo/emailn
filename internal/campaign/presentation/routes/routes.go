package routes

import (
	"emailn/cmd/api/handlers"
	"emailn/cmd/container/injector"
	"emailn/internal/campaign/presentation"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() http.Handler {
	r := chi.NewRouter()

	handler := injector.Resolve[presentation.Handler]()

	r.Post("/", handlers.ErrorHandler(handler.CreateCampaign))
	r.Get("/", handlers.ErrorHandler(handler.GetAll))
	r.Get("/{campaignId}", handlers.ErrorHandler(handler.GetCampaignById))

	return r
}
