package di

import (
	"github.com/gami/layered_arch_example/domain"
	"github.com/gami/layered_arch_example/repository"
)

func InjectTx() domain.Tx {
	return repository.NewDBTx(
		InjectUserDB(),
	)
}
