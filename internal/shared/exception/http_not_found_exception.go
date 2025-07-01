package exception

import "fmt"

type HttpNotFoundException struct {
	Message string
}

func (h HttpNotFoundException) Error() string {
	return fmt.Sprintf("Not Found Exception: %s", h.Message)
}

func NewHttpNotFoundException(message string) HttpNotFoundException {
	return HttpNotFoundException{
		Message: message,
	}
}
