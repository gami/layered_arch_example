package di

import (
	"app/domain/user"
	"app/messenger"
)

func InjectUserMessenger() user.Messenger {
	return messenger.NewUser(
		InjectSQS(),
	)
}
