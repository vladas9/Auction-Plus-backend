package repository

import (
	"database/sql"
	p "github.com/vladas9/backend-practice/pkg/postgres"
)

type StoreTx struct {
	*sql.Tx
}

func beginTx() (*StoreTx, error) {
	tx, err := p.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &StoreTx{tx}, nil
}

func WithTx(fn func(stx *StoreTx) error) error {
	storeTx, err := beginTx()
	defer storeTx.Rollback()
	if err != nil {
		return err
	}
	if err := fn(storeTx); err != nil {
		return err
	}
	return storeTx.Commit()
}
