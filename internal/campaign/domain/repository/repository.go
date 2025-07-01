package repository

import (
	"emailn/internal/campaign/domain/entity"
	sharedvo "emailn/internal/shared/valueobject"
)

type Repository interface {
	Create(entity *entity.Campaign) (*sharedvo.UID, error)
	GetAll() ([]*entity.Campaign, error)
	GetCampaignById(campaignId string) (*entity.Campaign, error)
}
