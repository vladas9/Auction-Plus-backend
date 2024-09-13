package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func TestBidRepo_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	bidRepo := NewBidRepo(tx)

	expectedBid := &u.GenerateDummyBids()[0]

	rows := sqlmock.NewRows([]string{"id", "user_id", "amount", "timestamp"}).
		AddRow(expectedBid.ID, expectedBid.UserId, expectedBid.Amount, expectedBid.Timestamp)

	mock.ExpectQuery(`
		SELECT 
			id,
			user_id,
			amount,
			timestamp
		FROM
			bids
		WHERE
			id = \$1`).
		WithArgs(expectedBid.Id()).
		WillReturnRows(rows)

	result, err := bidRepo.GetById(expectedBid.Id())
	assert.NoError(t, err)
	assert.Equal(t, expectedBid, result)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestBidRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	bidRepo := NewBidRepo(tx)

	bids := u.GenerateDummyBids()
	bid1 := &bids[1]
	bid2 := &bids[2]

	rows := sqlmock.NewRows([]string{"id", "user_id", "amount", "timestamp"}).
		AddRow(bid1.ID, bid1.UserId, bid1.Amount, bid1.Timestamp).
		AddRow(bid2.ID, bid2.UserId, bid2.Amount, bid2.Timestamp)

	mock.ExpectQuery(`
		SELECT 
			id,
			user_id,
			amount,
			timestamp
		FROM
			bids`).
		WillReturnRows(rows)

	results, err := bidRepo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, bid1, results[0])
	assert.Equal(t, bid2, results[1])

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestBidRepo_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	bidRepo := NewBidRepo(tx)

	newBid := &u.GenerateDummyBids()[1]

	mock.ExpectExec(`
		INSERT INTO bids \(
			id,
			user_id,
			amount,
			timestamp
		\) VALUES \(
			\$1, \$2, \$3, \$4
		\)
	`).WithArgs(
		newBid.Id(),
		newBid.UserId,
		newBid.Amount,
		newBid.Timestamp).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = bidRepo.Insert(newBid)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestBidRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	bidRepo := NewBidRepo(tx)

	updatedBid := &u.GenerateDummyBids()[2]

	mock.ExpectExec(`
		UPDATE 
			bids
		SET
			user_id = \$1,
			amount = \$2,
			timestamp = \$3
		WHERE
			id = \$4`).
		WithArgs(
			updatedBid.UserId,
			updatedBid.Amount,
			updatedBid.Timestamp,
			updatedBid.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = bidRepo.Update(updatedBid)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestBidRepo_Remove(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	bidRepo := NewBidRepo(tx)

	bidID := uuid.New()

	mock.ExpectExec(`
		DELETE FROM 
			bids
		WHERE 
			id = \$1`).
		WithArgs(bidID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = bidRepo.Remove(bidID)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}
