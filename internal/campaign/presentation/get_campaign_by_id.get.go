package presentation

import (
	"emailn/internal/campaign/presentation/dto"
	"emailn/internal/shared/exception"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) GetCampaignById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaignId := chi.URLParam(r, "campaignId")
	campaign, err := h.Service.GetCampaignById(campaignId)
	if err != nil {
		return nil, http.StatusInternalServerError, exception.NewHttpInternalServerException("Error getting campaign with ID " + campaignId)
	}

	if campaign == nil {
		return nil, http.StatusNotFound, exception.NewHttpNotFoundException("No campaign was found with the ID " + campaignId)
	}

	campaignDtoResponse := dto.NewGetCampaignDto(campaign)

	return campaignDtoResponse, http.StatusOK, nil
}
