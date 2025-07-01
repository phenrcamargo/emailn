package entity

import (
	"emailn/internal/campaign/application/contract"
	"emailn/internal/shared/exception"
	sharedvo "emailn/internal/shared/valueobject"
	"time"
)

type Contact struct {
	ID        *sharedvo.UID
	FirstName *sharedvo.Name
	LastName  *sharedvo.Name
	Phone     *sharedvo.Phone
	Email     *sharedvo.Email
	CreatedAt *sharedvo.DateTime
	UpdatedAt *sharedvo.DateTime
	DeletedAt *sharedvo.DateTime
}

func NewContact(input contract.NewContact) (*Contact, error) {
	if err := validateContactInput(input); err != nil {
		return nil, err
	}

	uid, err := sharedvo.NewUID()
	if err != nil {
		return nil, err
	}

	firstName, err := sharedvo.NewName(input.FirstName)
	if err != nil {
		return nil, err
	}

	lastName, err := sharedvo.NewName(input.LastName)
	if err != nil {
		return nil, err
	}

	phone, err := sharedvo.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	email, err := sharedvo.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	dateTime, err := sharedvo.NewDateTime(time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	contact := Contact{
		ID:        uid,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		CreatedAt: dateTime,
	}

	return &contact, nil

}

func validateContactInput(input contract.NewContact) error {
	if input.FirstName == "" {
		return exception.NewDomainException("First Name must not be empty")
	}

	if input.LastName == "" {
		return exception.NewDomainException("Last Name must not be empty")
	}

	if input.Phone == "" {
		return exception.NewDomainException("Phone Number must not be empty")
	}

	if input.Email == "" {
		return exception.NewDomainException("E-Mail Address must not be empty")
	}

	return nil
}
