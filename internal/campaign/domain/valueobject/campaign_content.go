package valueobject

import (
	"emailn/internal/shared/exception"
)

type CampaignContent string

func NewCampaignContent(s string) (*CampaignContent, error) {
	if len(s) > 255 {
		return nil, exception.NewDomainException("Content cannot be more than 255")
	}

	content := CampaignContent(s)
	return &content, nil
}

func (c CampaignContent) IsEqualContent(otherContent string) bool {
	return string(c) == otherContent
}
