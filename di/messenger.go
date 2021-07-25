package di

import (
	"app/domain/user"
	"app/infra/messenger"
)

func InjectUserMessenger() user.Messenger {
	return messenger.NewUser(
		InjectSQS(),
	)
}
