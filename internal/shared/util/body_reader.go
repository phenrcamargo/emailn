package util

import (
	"emailn/internal/shared/exception"
	"encoding/json"
	"io"
)

func BodyReader[T interface{}](body io.ReadCloser) (*T, error) {
	var data []byte
	var buf []byte

	buf = make([]byte, 1024)

	for {
		n, err := body.Read(buf)
		if n > 0 {
			data = append(data, buf[:n]...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, exception.NewHttpInternalServerException("Error when reading the request body")
		}
	}

	var bodyRequest T
	err := json.Unmarshal(data, &bodyRequest)
	if err != nil {
		return nil, exception.NewHttpInternalServerException("Error when decode body to json")
	}

	return &bodyRequest, nil
}
