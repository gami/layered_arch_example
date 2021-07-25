package profile

import (
	"errors"

	"app/domain/user"
)

type ID uint64

type Profile struct {
	ID     uint64
	UserID user.ID
	Hobby  string
}

func (p *Profile) Validate() error {
	if p.UserID < 1 {
		return errors.New("profile.user_id must not be zero")
	}

	if len(p.Hobby) > 1000 {
		return errors.New("profile.hobby is too long")
	}

	return nil
}
