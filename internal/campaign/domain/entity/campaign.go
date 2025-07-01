package entity

import (
	"emailn/internal/campaign/application/contract"
	"emailn/internal/campaign/domain/enum"
	campaignvo "emailn/internal/campaign/domain/valueobject"
	"emailn/internal/shared/exception"
	sharedvo "emailn/internal/shared/valueobject"
	"time"
)

type Campaign struct {
	ID        *sharedvo.UID
	Name      *sharedvo.Name
	Content   *campaignvo.CampaignContent
	Status    enum.CampaignStatus
	Contacts  []*Contact
	CreatedAt *sharedvo.DateTime
	UpdatedAt *sharedvo.DateTime
	DeletedAt *sharedvo.DateTime
}

func NewCampaign(input contract.NewCampaign) (*Campaign, error) {
	if err := validateCampaignInput(input); err != nil {
		return nil, err
	}

	uid, err := sharedvo.NewUID()
	if err != nil {
		return nil, err
	}

	name, err := sharedvo.NewName(input.Name)
	if err != nil {
		return nil, err
	}

	content, err := campaignvo.NewCampaignContent(input.Content)
	if err != nil {
		return nil, err
	}

	dateTime, err := sharedvo.NewDateTime(time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	contacts, err := buildContactsList(input.Contacts)
	if err != nil {
		return nil, err
	}

	return &Campaign{
		ID:        uid,
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		Status:    enum.Pending,
		CreatedAt: dateTime,
	}, nil
}

func validateCampaignInput(input contract.NewCampaign) error {
	if input.Name == "" {
		return exception.NewDomainException("Name must not be empty")
	}

	if input.Content == "" {
		return exception.NewDomainException("Content must not be empty")
	}

	if len(input.Contacts) <= 0 {
		return exception.NewDomainException("Contact list must not be empty")
	}

	return nil
}

func buildContactsList(contacts []contract.NewContact) ([]*Contact, error) {
	contactList := make([]*Contact, 0, len(contacts))

	for _, contact := range contacts {
		contactEntity, err := NewContact(contact)
		if err != nil {
			return nil, err
		}

		contactList = append(contactList, contactEntity)
	}

	return contactList, nil
}
