package valueobject

import (
	"emailn/internal/shared/exception"

	"github.com/rs/xid"
)

type UID string

func NewUID() (*UID, error) {
	uid := xid.New().String()

	if uid == "" {
		return nil, exception.NewDomainException("ID must not be empty")
	}

	guid := UID(uid)
	return &guid, nil

}

func (u UID) IsEqual(uid string) bool {
	return string(u) == uid
}
