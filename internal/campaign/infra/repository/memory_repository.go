package repository

import (
	"emailn/internal/campaign/domain/entity"
	sharedvo "emailn/internal/shared/valueobject"
)

type MemoryRepository struct {
	savedCampaigns []*entity.Campaign
}

func (mr *MemoryRepository) Create(entity *entity.Campaign) (*sharedvo.UID, error) {
	mr.savedCampaigns = append(mr.savedCampaigns, entity)
	return entity.ID, nil
}

func (mr *MemoryRepository) GetAll() ([]*entity.Campaign, error) {
	return mr.savedCampaigns, nil
}

func (mr *MemoryRepository) GetCampaignById(campaignId string) (*entity.Campaign, error) {
	for _, campaign := range mr.savedCampaigns {
		if campaign.ID.IsEqual(campaignId) {
			return campaign, nil
		}
	}

	return nil, nil
}
