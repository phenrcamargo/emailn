package valueobject

import "emailn/internal/shared/validator"

type DateTime string

func NewDateTime(s string) (*DateTime, error) {
	if err := validator.ValidaDateTime(s); err != nil {
		return nil, err
	}

	dateTime := DateTime(s)

	return &dateTime, nil
}
