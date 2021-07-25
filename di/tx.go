package di

import (
	"app/domain"
	"app/repository"
)

func InjectTx() domain.Tx {
	return repository.NewDBTx(
		InjectUserDB(),
	)
}
