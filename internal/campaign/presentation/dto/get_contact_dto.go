package dto

import (
	"emailn/internal/campaign/domain/entity"
	sharedvo "emailn/internal/shared/valueobject"
)

type GetContactDto struct {
	FirstName *sharedvo.Name
	LastName  *sharedvo.Name
	Phone     *sharedvo.Phone
	Email     *sharedvo.Email
}

func NewGetContactDto(contact *entity.Contact) GetContactDto {
	return GetContactDto{
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}
}

func NewGetContactDtoList(contacts []*entity.Contact) []GetContactDto {
	getContactDtoList := make([]GetContactDto, 0, len(contacts))

	for _, contactEntity := range contacts {
		getContactDtoList = append(getContactDtoList, NewGetContactDto(contactEntity))
	}

	return getContactDtoList
}
