package unitofwork

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	//r "github.com/vladas9/backend-practice/internal/repository"

	"testing"
	//"database/sql"
)

func TestNewUnitOfWork(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	uow := NewUnitOfWork(db)
	assert.NotNil(t, uow)

	err = uow.BeginTransaction()
	assert.NoError(t, err)

	assert.NotNil(t, uow.UserRepo)

	mock.ExpectCommit()
	err = uow.Commit()
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUnitOfWorkRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	uow := NewUnitOfWork(db)
	err = uow.BeginTransaction()
	assert.NoError(t, err)

	mock.ExpectRollback()
	err = uow.Rollback()
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
