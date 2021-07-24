package domain

import (
	"context"
)

// Tx is a interface of Transaction for Database.
type Tx interface {
	Transact(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
