package user

import "context"

type Messenger interface {
	SendCreated(ctx context.Context, msg *User) error
	RecieveCreated(ctx context.Context, f func(msg *User) error) error
}
