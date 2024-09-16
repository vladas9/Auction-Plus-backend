package repository

import (
	"database/sql"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type userRepo struct {
	tx *sql.Tx
}

func (s StoreTx) UserRepo() *userRepo {
	return &userRepo{s.Tx}
}

func (r *userRepo) GetById(id uuid.UUID) (*m.UserModel, error) {
	item := &m.UserModel{}
	query := `
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
			id = $1
	`
	row := r.tx.QueryRow(query, id)
	if err := row.Scan(
		&item.ID,
		&item.Username,
		&item.Email,
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

func (r *userRepo) Insert(item *m.UserModel) error {
	query := `
        INSERT INTO users (
            username,
            email,
            address,
            password,
            phone_number,
            user_type
        ) VALUES (
            $1, $2, $3, $4, $5, $6
        )
    `
	_, err := r.tx.Exec(query,
		item.Username,
		item.Email,
		item.Address,
		item.Password,
		item.PhoneNumber,
		item.UserType,
	)
	if err != nil {
		return err
	}

	return nil
}
