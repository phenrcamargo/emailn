package dto

import (
	"emailn/internal/campaign/domain/entity"
	campaignvo "emailn/internal/campaign/domain/valueobject"
	sharedvo "emailn/internal/shared/valueobject"
)

type GetCampaignDto struct {
	Name      sharedvo.Name
	CreatedOn sharedvo.DateTime
	Content   campaignvo.CampaignContent
	Contacts  []GetContactDto
	Status    string
}

func NewGetCampaignDto(campaign *entity.Campaign) GetCampaignDto {
	return GetCampaignDto{
		Name:      *campaign.Name,
		CreatedOn: *campaign.CreatedAt,
		Content:   *campaign.Content,
		Contacts:  NewGetContactDtoList(campaign.Contacts),
		Status:    campaign.Status.GetDescribe(),
	}
}

func NewGetCampaignDtoList(campaigns []*entity.Campaign) []GetCampaignDto {
	getCampaignDtoList := make([]GetCampaignDto, len(campaigns))

	for i, campaign := range campaigns {
		getCampaignDtoList[i] = NewGetCampaignDto(campaign)
	}

	return getCampaignDtoList
}
