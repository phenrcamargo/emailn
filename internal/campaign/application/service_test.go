package application

import (
	"emailn/internal/campaign/application/contract"
	"emailn/internal/campaign/domain/entity"
	"emailn/internal/shared/validator"
	"emailn/internal/shared/valueobject"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Create(entity *entity.Campaign) (*valueobject.UID, error) {
	args := r.Called(entity)

	uid, ok := args.Get(0).(*valueobject.UID)

	if !ok {
		return nil, fmt.Errorf("invalid return UID")
	}

	return uid, args.Error(1)
}

func (r *RepositoryMock) GetAll() ([]*entity.Campaign, error) {
	args := r.Called()

	return args.Get(0).([]*entity.Campaign), args.Error(1)
}

func (r *RepositoryMock) GetCampaignById(campaignId string) (*entity.Campaign, error) {
	args := r.Called(campaignId)

	return args.Get(0).(*entity.Campaign), args.Error(1)
}

var (
	repositoryMock = new(RepositoryMock)
	service        = Service{Repository: repositoryMock}
	expectedUID, _ = valueobject.NewUID()
)

func clearStubs() {
	repositoryMock.ExpectedCalls = nil
	repositoryMock.Calls = nil
}

func setupAll() {
	validator.ConfigValidator()
}

func TestMain(m *testing.M) {
	setupAll()
	m.Run()
}

func Test_Create_Campaign(t *testing.T) {
	/// Arrange
	clearStubs()
	assert := assert.New(t)

	newCampaign := contract.NewCampaign{
		Name:    "Test",
		Content: "Test Content",
		Contacts: []contract.NewContact{
			{
				FirstName: "Teste",
				LastName:  "Teste",
				Phone:     "999999999",
				Email:     "teste@email.com",
			},
		},
	}

	repositoryMock.
		On("Create", mock.MatchedBy(func(campaign *entity.Campaign) bool {
			if !campaign.Name.IsEqualNames(newCampaign.Name) {
				return false
			}

			if !campaign.Content.IsEqualContent(newCampaign.Content) {
				return false
			}

			if len(campaign.Contacts) != len(newCampaign.Contacts) {
				return false
			}

			return true
		})).
		Return(expectedUID, nil)

	/// Act
	campaignId, err := service.Create(newCampaign)

	/// Assert
	assert.Nil(err)
	assert.EqualValues(expectedUID, campaignId)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_Campaign_Domain_Error(t *testing.T) {
	/// Arrange
	clearStubs()
	assert := assert.New(t)

	invalidNewCampaign := contract.NewCampaign{
		Name:    "",
		Content: "Content testing",
		Contacts: []contract.NewContact{
			{
				FirstName: "Teste",
				LastName:  "Teste",
				Phone:     "999999999",
				Email:     "teste@email.com",
			},
			{
				FirstName: "Teste",
				LastName:  "Teste",
				Phone:     "999999999",
				Email:     "teste2@email.com",
			},
		},
	}

	/// Act
	_, err := service.Create(invalidNewCampaign)

	/// Assert
	assert.NotNil(err)
	assert.Contains(err.Error(), "Name must not be empty")
}

func Test_Create_Campaign_Repository_Error(t *testing.T) {
	/// Arrange
	clearStubs()
	assert := assert.New(t)

	newCampaign := contract.NewCampaign{
		Name:    "Test",
		Content: "Test Content",
		Contacts: []contract.NewContact{
			{
				FirstName: "Teste",
				LastName:  "Teste",
				Phone:     "999999999",
				Email:     "teste@email.com",
			},
		},
	}

	repositoryMock.
		On("Create", mock.Anything).
		Return(expectedUID, errors.New("error to save on database"))

	/// Act
	_, err := service.Create(newCampaign)

	/// Assert
	assert.NotNil(err)
	assert.Contains(err.Error(), "error to save on database")
	repositoryMock.AssertExpectations(t)
}
