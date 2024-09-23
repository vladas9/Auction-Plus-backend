package repository

import (
	"database/sql"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type bidRepo struct {
	tx *sql.Tx
}

func (s StoreTx) BidRepo() *bidRepo {
	return &bidRepo{s.Tx}
}

func (r *bidRepo) GetById(id uuid.UUID) (*m.BidModel, error) {
	item := &m.BidModel{}
	query := `
		SELECT 
			id,
			auction_id,
			bidder_id,
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
		&item.AuctionId,
		&item.UserId,
		&item.Amount,
		&item.Timestamp,
	); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *bidRepo) GetAllByUserId(id uuid.UUID, limit, offset int) ([]*m.BidModel, error) {
	var bids []*m.BidModel
	query := `
		SELECT 
			id,
			auction_id,
			amount,
			timestamp
		FROM
			bids
		WHERE
			bidder_id = $1
		LIMIT
			$2
		OFFSET
			$3
	`
	rows, err := r.tx.Query(query, id, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		item := &m.BidModel{}
		if err := rows.Scan(
			&item.ID,
			&item.AuctionId,
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

func (r *bidRepo) GetAllFor(auct *m.AuctionModel) ([]*m.BidModel, error) {
	var bids []*m.BidModel
	query := `
		SELECT 
			id,
			auction_id,
			bidder_id,
			amount,
			timestamp
		FROM
			bids
		WHERE auction_id = $1
			
	`
	rows, err := r.tx.Query(query, auct.Id())
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		item := &m.BidModel{}
		if err := rows.Scan(
			&item.ID,
			&item.AuctionId,
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

func (r *bidRepo) Insert(item *m.BidModel) error {
	query := `
		INSERT INTO bids (
			auction_id,
			bidder_id,
			amount
		) VALUES (
			$1, $2, $3
		)
	`
	_, err := r.tx.Exec(query,
		&item.AuctionId,
		&item.UserId,
		&item.Amount,
	)

	return err
}

func (r *bidRepo) Update(item *m.BidModel) error {
	query := `
		UPDATE 
			bids
		SET
			bidder_id = $1,
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
