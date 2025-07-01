package valueobject

import "emailn/internal/shared/validator"

type Email string

func NewEmail(s string) (*Email, error) {
	if err := validator.ValidateEmail(s); err != nil {
		return nil, err
	}

	email := Email(s)
	return &email, nil
}
