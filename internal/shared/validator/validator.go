package validator

import (
	"emailn/internal/shared/exception"
	"regexp"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func ConfigValidator() {
	validate.RegisterValidation("datetime", func(fl validator.FieldLevel) bool {
		dateTimeIso8601Regex := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:Z|[+-]\d{2}:\d{2})?$`
		regexp := regexp.MustCompile(dateTimeIso8601Regex)

		return regexp.MatchString(fl.Field().String())
	})

	validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phoneRegex := `^\+?[1-9][0-9]{7,14}$`
		regexp := regexp.MustCompile(phoneRegex)

		return regexp.MatchString(fl.Field().String())
	})
}

func ValidateEmail(e string) error {
	if err := validate.Var(e, "required,email"); err == nil {
		return nil
	}

	return exception.NewDomainException("Invalid email")
}

func ValidateName(e string) error {
	if err := validate.Var(e, "required,max=50"); err == nil {
		return nil
	}

	return exception.NewDomainException("Invalid name")
}

func ValidaDateTime(d string) error {
	if err := validate.Var(d, "required,datetime"); err == nil {
		return nil
	}

	return exception.NewDomainException("Invalid date time")
}

func ValidatePhone(p string) error {
	if err := validate.Var(p, "required,phone"); err == nil {
		return nil
	}

	return exception.NewDomainException("Invalid phone number")
}
