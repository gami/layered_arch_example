package di

import (
	"app/domain"
	"app/infra/repository"
)

func InjectTx() domain.Tx {
	return repository.NewDBTx(
		InjectUserDB(),
	)
}
