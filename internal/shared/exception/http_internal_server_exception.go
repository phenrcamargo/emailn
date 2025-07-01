package exception

import "fmt"

type HttpInternalServerException struct {
	Message string
}

func (h HttpInternalServerException) Error() string {
	return fmt.Sprintf("Internal Server Error: %s", h.Message)
}

func NewHttpInternalServerException(message string) HttpInternalServerException {
	return HttpInternalServerException{
		Message: message,
	}
}
