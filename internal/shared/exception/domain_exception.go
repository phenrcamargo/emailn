package exception

import "fmt"

type DomainException struct {
	Message string
}

func (e DomainException) Error() string {
	return fmt.Sprintf("Domain Exception: %s", e.Message)
}

func NewDomainException(message string) DomainException {
	return DomainException{
		Message: message,
	}
}
