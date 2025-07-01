package entity

import (
	"emailn/internal/campaign/application/contract"
	"emailn/internal/campaign/domain/enum"
	"emailn/internal/shared/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupAll() {
	validator.ConfigValidator()
}

func TestMain(m *testing.M) {
	setupAll()
	m.Run()
}

func Test_Create_Campaign_Success(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	input := contract.NewCampaign{
		Name:    "Test",
		Content: "Content testing",
		Contacts: []contract.NewContact{
			{
				FirstName: "Teste",
				LastName:  "Teste",
				Phone:     "999999999",
				Email:     "teste@email.com",
			},
		},
	}

	/// Act
	campaign, err := NewCampaign(input)

	/// Assert
	assert.Nil(err)
	assert.NotNil(campaign)
	assert.NotNil(campaign.ID)
	assert.EqualValues(*campaign.Name, input.Name)
	assert.EqualValues(*campaign.Content, input.Content)
	assert.Equal(len(campaign.Contacts), len(input.Contacts))
	assert.Equal(campaign.Status, enum.Pending)
}

func Test_Create_Campaign_Name_Error(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	invalidInput := contract.NewCampaign{
		Name:    "",
		Content: "Content testing",
		Contacts: []contract.NewContact{
			{
				FirstName: "Teste 2",
				LastName:  "Teste 2",
				Phone:     "999999999",
				Email:     "teste2@email.com",
			},
		},
	}

	/// Act
	campaign, err := NewCampaign(invalidInput)

	/// Assert
	assert.Nil(campaign)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Name must not be empty")
}

func Test_Create_Campaign_Content_Error(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	invalidInput := contract.NewCampaign{
		Name:    "Test",
		Content: "",
		Contacts: []contract.NewContact{
			{
				FirstName: "Teste",
				LastName:  "Teste",
				Phone:     "999999999",
				Email:     "teste@email.com",
			},
			{
				FirstName: "Teste 2",
				LastName:  "Teste 2",
				Phone:     "999999999",
				Email:     "teste2@email.com",
			},
		},
	}

	/// Act
	campaign, err := NewCampaign(invalidInput)

	/// Assert
	assert.Nil(campaign)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Content must not be empty")
}

func Test_Create_Campaign_Contacts_Error(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	invalidInput := contract.NewCampaign{
		Name:     "Test",
		Content:  "Content testing",
		Contacts: []contract.NewContact{},
	}

	/// Act
	campaign, err := NewCampaign(invalidInput)

	/// Assert
	assert.Nil(campaign)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Contact list must not be empty")
}

func Test_Create_Campaign_InvalidEmail_Error(t *testing.T) {
	/// Arrange
	assert := assert.New(t)
	invalidInput := contract.NewCampaign{
		Name:    "Test",
		Content: "Content testing",
		Contacts: []contract.NewContact{
			{
				FirstName: "Teste",
				LastName:  "Teste",
				Phone:     "999999999",
				Email:     "invalidemail",
			},
		},
	}

	/// Act
	campaign, err := NewCampaign(invalidInput)

	/// Assert
	assert.Nil(campaign)
	assert.NotNil(err)
	assert.Contains(err.Error(), "Invalid email")
}
