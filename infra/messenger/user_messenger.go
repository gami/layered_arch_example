package messenger

import (
	"context"

	"app/adapter/aws"
	"app/domain/user"
)

const keyUserCreated = "user_created"

type User struct {
	sqs *aws.SQS
}

var _ user.Messenger = &User{}

func NewUser(sqs *aws.SQS) *User {
	return &User{
		sqs: sqs,
	}
}

func (m *User) SendCreated(ctx context.Context, u *user.User) error {
	data, err := u.Marshal()
	if err != nil {
		return err
	}

	return m.sqs.Send(ctx, keyUserCreated, data)
}

func (m *User) RecieveCreated(ctx context.Context, f func(msg *user.User) error) error {
	return m.sqs.Recieve(ctx, keyUserCreated, func(data string) error {
		u := &user.User{}

		if err := u.Unmarshal(data); err != nil {
			return err
		}

		return f(u)
	})
}
