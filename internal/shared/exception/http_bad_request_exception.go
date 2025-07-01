package exception

import "fmt"

type HttpBadRequestException struct {
	Message string
}

func (h HttpBadRequestException) Error() string {
	return fmt.Sprintf("Bad Request: %s", h.Message)
}

func NewHttpBadRequestException(message string) HttpBadRequestException {
	return HttpBadRequestException{
		Message: message,
	}
}
