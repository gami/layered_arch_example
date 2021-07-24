package profile

import "github.com/gami/layered_arch_example/domain/user"

type ID uint64

type Profile struct {
	ID     uint64
	UserID user.ID
	Hobby  string
}

func (p *Profile) Validate() error {
	return nil
}
