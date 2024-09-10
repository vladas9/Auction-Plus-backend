package repository

import (
	"database/sql"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type AuctionRepo struct {
	tx *sql.Tx
}

func NewAuctionRepo(tx *sql.Tx) *AuctionRepo {
	return &AuctionRepo{tx}
}

func (r *AuctionRepo) GetById(id uuid.UUID) (*m.AuctionModel, error) {
	item := &m.AuctionModel{}
	query := `
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
			id = $1`
	row := r.tx.QueryRow(query, id)
	if err := row.Scan(
		&item.ID,
		&item.StartingBid,
		&item.ClosingBid,
		&item.StartTime,
		&item.EndTime,
		&item.ExtraTimeEnabled,
		&item.ExtraTimeDuration,
		&item.ExtraTimeThreshold,
		&item.IsActive,
	); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *AuctionRepo) GetAll() ([]*m.AuctionModel, error) {
	var auctions []*m.AuctionModel
	query := `
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
	`
	rows, err := r.tx.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		item := &m.AuctionModel{}
		if err := rows.Scan(
			&item.ID,
			&item.SellerId,
			&item.StartingBid,
			&item.ClosingBid,
			&item.StartTime,
			&item.EndTime,
			&item.ExtraTimeEnabled,
			&item.ExtraTimeDuration,
			&item.ExtraTimeThreshold,
			&item.IsActive,
		); err != nil {
			return nil, err
		}
		auctions = append(auctions, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return auctions, nil
}

func (r *AuctionRepo) Update(item *m.AuctionModel) error {
	query := `
		UPDATE 
			auctions
		SET
			seller_id = $1,
			starting_bid = $2,
			closing_bid = $3,
			start_time = $4,
			end_time = $5,
			extra_time_enabled = $6,
			extra_time_duration = $7,
			extra_time_threshold = $8,
			is_active = $9
		WHERE
			id = $10
	`
	_, err := r.tx.Exec(query,
		item.SellerId,
		item.StartingBid,
		item.ClosingBid,
		item.StartTime,
		item.EndTime,
		item.ExtraTimeEnabled,
		item.ExtraTimeDuration,
		item.ExtraTimeThreshold,
		item.IsActive,
		item.ID,
	)

	return err
}

func (r *AuctionRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			auctions
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

func (r *AuctionRepo) Insert(item *m.AuctionModel) error {
	query := `
		INSERT INTO auctions (
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
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)
	`
	_, err := r.tx.Exec(query,
		item.ID,
		item.SellerId,
		item.StartingBid,
		item.ClosingBid,
		item.StartTime,
		item.EndTime,
		item.ExtraTimeEnabled,
		item.ExtraTimeDuration,
		item.ExtraTimeThreshold,
		item.IsActive,
	)

	return err
}
