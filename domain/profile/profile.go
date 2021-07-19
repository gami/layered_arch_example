package profile

type Profile struct {
	ID     uint64
	UserID uint64
	Hobby  string
}

func (p *Profile) Validate() error {
	return nil
}
