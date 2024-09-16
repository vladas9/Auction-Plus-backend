package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vladas9/backend-practice/internal/utils"
	"testing"
)

func TestUserRepo_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	userRepo := NewUserRepo(tx)

	expectedUser := utils.GenerateDummyUser()
	rows := sqlmock.NewRows(
		[]string{
			"id",
			"username",
			"email",
			"password",
			"address",
			"phone_number",
			"user_type",
			"registered_date"}).
		AddRow(expectedUser.Id(),
			expectedUser.Username,
			expectedUser.Email,
			expectedUser.Password,
			expectedUser.Address,
			expectedUser.PhoneNumber,
			expectedUser.UserType, expectedUser.RegisteredDate)

	mock.ExpectQuery(`
		SELECT
			id,
			username,
			email,
			password,
			address,
			phone_number,
			user_type,
			registered_date
		FROM
			users
		WHERE
			id = \$1`).
		WithArgs(expectedUser.Id()).
		WillReturnRows(rows)

	result, err := userRepo.GetById(expectedUser.Id())
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestUserRepo_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	userRepo := NewUserRepo(tx)

	newUser := utils.GenerateDummyUser()
	mock.ExpectExec(`INSERT INTO users`).
		WithArgs(
			newUser.Id(),
			newUser.Username,
			newUser.Email,
			newUser.Password,
			newUser.Address,
			newUser.PhoneNumber,
			newUser.UserType,
			newUser.RegisteredDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = userRepo.Insert(newUser)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestUserRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	userRepo := NewUserRepo(tx)

	updatedUser := utils.GenerateDummyUser()
	mock.ExpectExec(`
		UPDATE 
			users
		SET
			username = \$1,
			email = \$2,
			address = \$3,
			password = \$4
			phone_number = \$5,
			user_type = \$6,
			registered_date = \$7
		WHERE
			id = \$8
	`).
		WithArgs(
			updatedUser.Username,
			updatedUser.Email,
			updatedUser.Address,
			updatedUser.Password,
			updatedUser.PhoneNumber,
			updatedUser.UserType,
			updatedUser.RegisteredDate,
			updatedUser.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.Update(updatedUser)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestUserRepo_Remove(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	userRepo := NewUserRepo(tx)

	userID := utils.GenerateDummyUser().Id()

	mock.ExpectExec(`DELETE FROM users WHERE id = \$1`).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepo.Remove(userID)
	assert.NoError(t, err)

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}

func TestUserRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	assert.NoError(t, err)

	userRepo := NewUserRepo(tx)

	user1 := utils.GenerateDummyUser()
	user2 := utils.GenerateDummyUser()

	rows := sqlmock.NewRows([]string{
		"id", "username", "email", "password", "address", "phone_number", "user_type", "registered_date",
	}).AddRow(
		user1.Id(), user1.Username, user1.Email, user1.Password, user1.Address, user1.PhoneNumber, user1.UserType, user1.RegisteredDate,
	).AddRow(
		user2.Id(), user2.Username, user2.Email, user2.Password, user2.Address, user2.PhoneNumber, user2.UserType, user2.RegisteredDate,
	)

	mock.ExpectQuery(`SELECT id, username, email, password, address, phone_number, user_type, registered_date FROM users`).
		WillReturnRows(rows)

	result, err := userRepo.GetAll()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, user1, result[0])
	assert.Equal(t, user2, result[1])

	mock.ExpectCommit()
	err = tx.Commit()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}
}
