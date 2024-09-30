package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	"github.com/vladas9/backend-practice/internal/utils"
)

type userRepo struct {
	tx *sql.Tx
}

func (s *StoreTx) UserRepo() *userRepo {
	return &userRepo{s.Tx}
}

func (r *userRepo) GetByProperty(field string, value interface{}) (*m.UserModel, error) {
	// Map to validate allowable fields and their expected data types
	validFields := map[string]string{
		"id":           "id",
		"email":        "email",
		"username":     "username",
		"phone_number": "phone_number",
		// Add any other fields that you want to allow filtering by
	}

	// Check if the field provided is valid
	fieldQuery, valid := validFields[field]
	if !valid {
		return nil, fmt.Errorf("invalid field: %s", field)
	}

	// Base query with a placeholder for dynamic field
	query := fmt.Sprintf(`
        SELECT 
            id,
            username,
            email,
            image,
            password,
            address,
            phone_number,
            user_type,
            registered_date
        FROM
            users
        WHERE
            %s = $1
    `, fieldQuery)

	// Struct to hold the result
	item := &m.UserModel{}

	// Execute the query with the value parameter
	row := r.tx.QueryRow(query, value)
	if err := row.Scan(
		&item.ID,
		&item.Username,
		&item.Email,
		&item.Image,
		&item.Password,
		&item.Address,
		&item.PhoneNumber,
		&item.UserType,
		&item.RegisteredDate,
	); err != nil {
		return nil, err
	}

	return item, nil
}

func (r *userRepo) GetAll() ([]*m.UserModel, error) {
	var users []*m.UserModel
	query := `
		SELECT 
			id,
			username,
			email,
			password,
			image,
			address,
			phone_number,
			user_type,
			registered_date
		FROM
			users
	`
	rows, err := r.tx.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		item := &m.UserModel{}
		if err := rows.Scan(
			&item.ID,
			&item.Username,
			&item.Email,
			&item.Password,
			&item.Image,
			&item.Address,
			&item.PhoneNumber,
			&item.UserType,
			&item.RegisteredDate,
		); err != nil {
			return nil, err
		}
		users = append(users, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Update(item *m.UserModel) error {
	query := `
		UPDATE 
			users
		SET
			username = $1,
			email = $2,
			address = $3,
			password = $4
			phone_number = $5,
			user_type = $6,
			registered_date = $7
		WHERE
			id = $8
	`
	_, err := r.tx.Exec(query,
		&item.Username,
		&item.Email,
		&item.Address,
		&item.Password,
		&item.PhoneNumber,
		&item.UserType,
		&item.RegisteredDate,
		&item.ID,
	)

	return err
}

func (r *userRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			users
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

func (r *userRepo) Insert(item *m.UserModel) (uuid.UUID, error) {
	query := `
        INSERT INTO users (
            username,
            email,
  					image,
            address,
            password,
            phone_number,
            user_type
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7
				) RETURNING id
    `
	var userId string

	err := r.tx.QueryRow(query,
		item.Username,
		item.Email,
		item.Image,
		item.Address,
		item.Password,
		item.PhoneNumber,
		item.UserType,
	).Scan(&userId)
	utils.Logger.Info("userRepo: returning id after insert::", userId)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.MustParse(userId), nil
}
