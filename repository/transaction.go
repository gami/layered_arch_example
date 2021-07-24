package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/gami/layered_arch_example/adapter/mysql"
)

type txKey string

const (
	txKeyTag txKey = "txKey:tag"
)

// DBTx is a implemented transaction for DB
type DBTx struct {
	dbs []*mysql.DB
}

// NewDBTx is a function to initialize a Database Transaction. This support multiple DBs.
func NewDBTx(dbs ...*mysql.DB) *DBTx {
	return &DBTx{
		dbs: dbs,
	}
}

// Transact is a function to execute a process with using database queries. If some errors occurred, this will do rollback
// the transaction. If succeeded, this will do commit. This function is allowed to take only one transaction at same time.
func (t *DBTx) Transact(ctx context.Context, process func(context.Context) (interface{}, error)) (interface{}, error) {
	if InTransaction(ctx) {
		log.Println("this context has already other transaction")

		return process(ctx)
	}

	txs := make([]*sql.Tx, 0, len(t.dbs))
	defer t.clear(ctx)
	for _, db := range t.dbs {
		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return nil, errors.Wrap(err, "begin transaction failed")
		}

		ctx = context.WithValue(ctx, txKeyTag, tx)

		txs = append(txs, tx)
	}

	obj, err := process(ctx)

	if err == nil {
		for _, tx := range txs {
			err = tx.Commit()
		}
	}

	if err != nil {
		errR := rollback(txs)
		if errR != nil {
			err = errors.Wrap(err, fmt.Sprintf("rollback err = %v", errR))
		}
		return nil, err
	}

	return obj, nil
}

// clear is a function to remove reference about *sql.Tx from given context.
func (t *DBTx) clear(ctx context.Context) {
	_ = context.WithValue(ctx, txKeyTag, nil)
}

func rollback(txs []*sql.Tx) error {
	var err error
	for _, tx := range txs {
		err = tx.Rollback()
	}
	if err != nil {
		return err
	}
	return nil
}

// GetTx returns *sql.Tx for TagDB from context.
func GetTx(ctx context.Context) (*sql.Tx, bool) {
	return getTx(ctx, txKeyTag)
}

// InTransaction returns if given context has transactions.
func InTransaction(ctx context.Context) bool {
	_, res := getTx(ctx, txKeyTag)
	return res
}

func getTx(ctx context.Context, key txKey) (*sql.Tx, bool) {
	v := ctx.Value(key)
	tx, ok := v.(*sql.Tx)
	if !ok {
		return nil, false
	}
	return tx, true
}
