package presentation

import (
	"emailn/internal/campaign/presentation/dto"
	"emailn/internal/shared/exception"
	"net/http"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.Service.GetAll()
	if err != nil {
		return nil, http.StatusInternalServerError, exception.NewHttpInternalServerException("Error getting campaigns")
	}

	if len(campaigns) <= 0 {
		return nil, http.StatusNotFound, exception.NewHttpNotFoundException("Campaigns not found")
	}

	campaignsDtoResponse := dto.NewGetCampaignDtoList(campaigns)

	return campaignsDtoResponse, http.StatusOK, nil
}
