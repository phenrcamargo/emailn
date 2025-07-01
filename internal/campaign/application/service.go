package application

import (
	"emailn/internal/campaign/application/contract"
	"emailn/internal/campaign/domain/entity"
	"emailn/internal/campaign/domain/repository"
	"emailn/internal/shared/exception"
	sharedvo "emailn/internal/shared/valueobject"
)

type Service struct {
	Repository repository.Repository
}

func (s Service) Create(newCampaign contract.NewCampaign) (*sharedvo.UID, error) {
	campaignEntity, err := entity.NewCampaign(newCampaign)
	if err != nil {
		return nil, exception.NewHttpBadRequestException("Error when creating campaign entity: " + err.Error())
	}

	campaignId, err := s.Repository.Create(campaignEntity)
	if err != nil {
		return nil, exception.NewHttpInternalServerException("Error when creating campaign in database: " + err.Error())
	}

	return campaignId, nil
}

func (s Service) GetAll() ([]*entity.Campaign, error) {
	campaigns, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (s Service) GetCampaignById(campaignId string) (*entity.Campaign, error) {
	campaign, err := s.Repository.GetCampaignById(campaignId)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
