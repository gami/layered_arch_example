package user

import (
	"bytes"
	"encoding/json"

	"github.com/friendsofgo/errors"
)

type ID uint64

type User struct {
	ID   ID
	Name string
}

func (u *User) Validate() error {
	return nil
}

func (u *User) Marshal() (serialized string, err error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", errors.Wrapf(err, "failed to marshal user id=%d", u.ID)
	}

	return string(b), nil
}

func (u *User) Unmarshal(data string) error {
	err := json.Unmarshal(bytes.NewBufferString(data).Bytes(), u)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal user")
	}

	return nil
}
