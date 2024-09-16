package repository

import (
	"database/sql"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type bidRepo struct {
	tx *sql.Tx
}

func BidRepo(tx *sql.Tx) *bidRepo {
	return &bidRepo{tx}
}

func (r *bidRepo) GetById(id uuid.UUID) (*m.BidModel, error) {
	item := &m.BidModel{}
	query := `
		SELECT 
			id,
			user_id,
			amount,
			timestamp
		FROM
			bids
		WHERE
			id = $1
	`
	row := r.tx.QueryRow(query, id)
	if err := row.Scan(
		&item.ID,
		&item.UserId,
		&item.Amount,
		&item.Timestamp,
	); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *bidRepo) GetAll() ([]*m.BidModel, error) {
	var bids []*m.BidModel
	query := `
		SELECT 
			id,
			user_id,
			amount,
			timestamp
		FROM
			bids
	`
	rows, err := r.tx.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		item := &m.BidModel{}
		if err := rows.Scan(
			&item.ID,
			&item.UserId,
			&item.Amount,
			&item.Timestamp,
		); err != nil {
			return nil, err
		}
		bids = append(bids, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bids, nil
}

func (r *bidRepo) Update(item *m.BidModel) error {
	query := `
		UPDATE 
			bids
		SET
			user_id = $1,
			amount = $2,
			timestamp = $3
		WHERE
			id = $4
	`
	_, err := r.tx.Exec(query,
		&item.UserId,
		&item.Amount,
		&item.Timestamp,
		&item.ID,
	)

	return err
}

func (r *bidRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			bids
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

func (r *bidRepo) Insert(item *m.BidModel) error {
	query := `
		INSERT INTO bids (
			id,
			user_id,
			amount,
			timestamp
		) VALUES (
			$1, $2, $3, $4
		)
	`
	_, err := r.tx.Exec(query,
		&item.ID,
		&item.UserId,
		&item.Amount,
		&item.Timestamp,
	)

	return err
}
