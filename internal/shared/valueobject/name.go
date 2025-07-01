package valueobject

import "emailn/internal/shared/validator"

type Name string

func NewName(s string) (*Name, error) {
	if err := validator.ValidateName(s); err != nil {
		return nil, err
	}

	name := Name(s)
	return &name, nil
}

func (n Name) IsEqualNames(otherName string) bool {
	return string(n) == otherName
}
