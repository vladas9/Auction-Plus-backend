package repository

import (
	"database/sql"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type BidRepo struct {
	tx *sql.Tx
}

func NewBidRepo(tx *sql.Tx) *BidRepo {
	return &BidRepo{tx}
}

func (r *BidRepo) GetById(id uuid.UUID) (*m.BidModel, error) {
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

func (r *BidRepo) GetAll() ([]*m.BidModel, error) {
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

func (r *BidRepo) Update(item *m.BidModel) error {
	query := `
		UPDATE 
			bids
		SET
			id = $1,
			user_id = $2,
			amount = $3
			timestamp = $4
		WHERE
			id = $5
	`
	_, err := r.tx.Exec(query,
		&item.ID,
		&item.UserId,
		&item.Amount,
		&item.Timestamp,
	)

	return err
}

func (r *BidRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			bids
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

func (r *BidRepo) Insert(item *m.BidModel) error {
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
