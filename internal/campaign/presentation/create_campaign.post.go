package presentation

import (
	"emailn/internal/campaign/application/contract"
	"emailn/internal/shared/util"
	"net/http"
)

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	newCampaing, err := util.BodyReader[contract.NewCampaign](r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	id, err := h.Service.Create(*newCampaing)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response := map[string]string{"ObjectID": string(*id)}

	return response, http.StatusCreated, nil
}
