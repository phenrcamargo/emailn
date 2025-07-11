package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaignContent_Success(t *testing.T) {
	campaignContent, err := NewCampaignContent("content")

	assert.Nil(t, err)
	assert.NotNil(t, campaignContent)
}

func Test_NewCampaignContent_Error(t *testing.T) {
	campaignContent, err := NewCampaignContent("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has sur")

	assert.Nil(t, campaignContent)
	assert.NotNil(t, err)
}
