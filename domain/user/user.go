package user

type User struct {
	ID uint64
}

func (u *User) Validate() error {
	return nil
}
