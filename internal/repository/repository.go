package repository

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

type StoreTx struct {
	*sql.Tx
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) BeginTx() (*StoreTx, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	return &StoreTx{tx}, nil
}

func (s *Store) WithTx(fn func(stx *StoreTx) error) error {
	storeTx, err := s.BeginTx()
	defer storeTx.Rollback()
	if err != nil {
		return err
	}
	if err := fn(storeTx); err != nil {
		return err
	}
	return storeTx.Commit()
}
