package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vladas9/backend-practice/internal/utils"
	"testing"
)

func TestAuctionRepo_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	auctionRepo := NewAuctionRepo(tx)

	expectedAuction := utils.GenerateDummyAuctions()[0]
	rows := mock.NewRows(
		[]string{
			"id",
			"seller_id",
			"starting_bid",
			"closing_bid",
			"start_time",
			"end_time",
			"extra_time_enabled",
			"extra_time_duration",
			"extra_time_threshold",
			"is_active"}).
		AddRow(
			expectedAuction.Id(),
			expectedAuction.SellerId,
			expectedAuction.StartingBid,
			expectedAuction.ClosingBid,
			expectedAuction.StartTime,
			expectedAuction.EndTime,
			expectedAuction.ExtraTimeEnabled,
			expectedAuction.ExtraTimeDuration,
			expectedAuction.ExtraTimeThreshold,
			expectedAuction.IsActive)

	mock.ExpectQuery(`
		SELECT
			id,
			seller_id,
			starting_bid,
			closing_bid,
			start_time,
			end_time,
			extra_time_enabled,
			extra_time_duration,
			extra_time_threshold,
			is_active
		FROM
			auctions
		WHERE
			id = \$1`).
		WithArgs(expectedAuction.Id()).
		WillReturnRows(rows)

	result, err := auctionRepo.GetById(expectedAuction.Id())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &expectedAuction, result)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestAuctionRepo_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	auctionRepo := NewAuctionRepo(tx)

	newAuction := &utils.GenerateDummyAuctions()[0]

	// Mock the expected insert query
	mock.ExpectExec(`INSERT INTO auctions`).
		WithArgs(
			newAuction.Id(),
			newAuction.SellerId,
			newAuction.StartingBid,
			newAuction.ClosingBid,
			newAuction.StartTime,
			newAuction.EndTime,
			newAuction.ExtraTimeEnabled,
			newAuction.ExtraTimeDuration,
			newAuction.ExtraTimeThreshold,
			newAuction.IsActive).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call Insert and check for errors
	err = auctionRepo.Insert(newAuction)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestAuctionRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	auctionRepo := NewAuctionRepo(tx)

	updatedAuction := &utils.GenerateDummyAuctions()[0]

	// Mock the update query
	mock.ExpectExec(`UPDATE auctions SET`).
		WithArgs(
			updatedAuction.SellerId,
			updatedAuction.StartingBid,
			updatedAuction.ClosingBid,
			updatedAuction.StartTime,
			updatedAuction.EndTime,
			updatedAuction.ExtraTimeEnabled,
			updatedAuction.ExtraTimeDuration,
			updatedAuction.ExtraTimeThreshold,
			updatedAuction.IsActive,
			updatedAuction.Id(),
		).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call Update and check for errors
	err = auctionRepo.Update(updatedAuction)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestAuctionRepo_Remove(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	auctionRepo := NewAuctionRepo(tx)

	auctionID := utils.GenerateDummyAuctions()[0].Id()

	// Mock the delete query
	mock.ExpectExec(`DELETE FROM auctions WHERE id = \$1`).
		WithArgs(auctionID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call Remove and check for errors
	err = auctionRepo.Remove(auctionID)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestAuctionRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	auctionRepo := NewAuctionRepo(tx)

	auctions := utils.GenerateDummyAuctions()

	// Mock the expected result set
	rows := mock.NewRows([]string{
		"id", "seller_id", "starting_bid", "closing_bid", "start_time", "end_time",
		"extra_time_enabled", "extra_time_duration", "extra_time_threshold", "is_active",
	}).AddRow(
		auctions[0].Id(),
		auctions[0].SellerId,
		auctions[0].StartingBid,
		auctions[0].ClosingBid,
		auctions[0].StartTime,
		auctions[0].EndTime,
		auctions[0].ExtraTimeEnabled,
		auctions[0].ExtraTimeDuration,
		auctions[0].ExtraTimeThreshold,
		auctions[0].IsActive,
	).AddRow(
		auctions[1].Id(),
		auctions[1].SellerId,
		auctions[1].StartingBid,
		auctions[1].ClosingBid,
		auctions[1].StartTime,
		auctions[1].EndTime,
		auctions[1].ExtraTimeEnabled,
		auctions[1].ExtraTimeDuration,
		auctions[1].ExtraTimeThreshold,
		auctions[1].IsActive,
	)

	mock.ExpectQuery(`SELECT id, seller_id, starting_bid, closing_bid, start_time, end_time, extra_time_enabled, extra_time_duration, extra_time_threshold, is_active FROM auctions`).
		WillReturnRows(rows)

	// Call the GetAll method
	result, err := auctionRepo.GetAll()
	assert.NoError(t, err)

	// Check the result
	assert.Equal(t, 2, len(result))
	assert.Equal(t, &auctions[0], result[0])
	assert.Equal(t, &auctions[1], result[1])

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}
