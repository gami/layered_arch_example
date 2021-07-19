package domain

import "context"

const (
	KeyUserCreated = "user_created"
)

type Message interface {
	Marshal() (serialized string, err error)
	Unmarshal(data string) error
}

type Messenger interface {
	Send(ctx context.Context, key string, msg Message) error
	RecieveUser(ctx context.Context, key string, f func(msg Message) error) error
}
