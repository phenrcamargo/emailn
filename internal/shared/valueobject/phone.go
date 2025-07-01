package valueobject

import "emailn/internal/shared/validator"

type Phone string

func NewPhone(s string) (*Phone, error) {
	if err := validator.ValidatePhone(s); err != nil {
		return nil, err
	}

	phone := Phone(s)
	return &phone, nil
}

func (n Phone) IsEqualPhones(otherPhone string) bool {
	return string(n) == otherPhone
}
