package repository

import (
	"database/sql"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type auctionRepo struct {
	tx *sql.Tx
}

func (s *StoreTx) AuctionRepo() *auctionRepo {
	return &auctionRepo{tx: s.Tx}
}

func (r *auctionRepo) GetById(id uuid.UUID) (*m.AuctionModel, error) {
	item := &m.AuctionModel{}
	query := `
		SELECT 
			id,
			seller_id,
			start_price,
			current_bid,
			start_time,
			end_time,
			extra_time_enabled,
			extra_time_duration,
			extra_time_threshold,
			status
		FROM
			auctions
		WHERE
			id = $1`
	row := r.tx.QueryRow(query, id)
	if err := row.Scan(
		&item.ID,
		&item.SellerId,
		&item.StartPrice,
		&item.CurrentBid,
		&item.StartTime,
		&item.EndTime,
		&item.ExtraTimeEnabled,
		&item.ExtraTimeDuration,
		&item.ExtraTimeThreshold,
		&item.Status,
	); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *auctionRepo) GetAll(limit, offset int) ([]*m.AuctionModel, error) {
	var auctions []*m.AuctionModel
	query := `
		SELECT 
			id,
			seller_id,
			start_price,
			current_bid,
			start_time,
			end_time,
			extra_time_enabled,
			extra_time_duration,
			extra_time_threshold,
			status
		FROM
			auctions
		ORDER BY
			start_time
		LIMIT
			$1
		OFFSET
			$2
	`
	rows, err := r.tx.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		item := &m.AuctionModel{}
		if err := rows.Scan(
			&item.ID,
			&item.SellerId,
			&item.StartPrice,
			&item.CurrentBid,
			&item.StartTime,
			&item.EndTime,
			&item.ExtraTimeEnabled,
			&item.ExtraTimeDuration,
			&item.ExtraTimeThreshold,
			&item.Status,
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

func (r *auctionRepo) Update(item *m.AuctionModel) error {
	query := `
		UPDATE 
			auctions
		SET
			seller_id = $1,
			start_price = $2,
			current_bid = $3,
			start_time = $4,
			end_time = $5,
			extra_time_enabled = $6,
			extra_time_duration = $7,
			extra_time_threshold = $8,
			status = $9
		WHERE
			id = $10
	`
	_, err := r.tx.Exec(query,
		item.SellerId,
		item.StartPrice,
		item.CurrentBid,
		item.StartTime,
		item.EndTime,
		item.ExtraTimeEnabled,
		item.ExtraTimeDuration,
		item.ExtraTimeThreshold,
		item.Status,
		item.ID,
	)

	return err
}

func (r *auctionRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			auctions
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

func (r *auctionRepo) Insert(item *m.AuctionModel) error {
	query := `
		INSERT INTO auctions (
			seller_id,
			start_price,
			current_bid,
			start_time,
			end_time,
			extra_time_enabled,
			extra_time_duration,
			extra_time_threshold,
			status
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
	`
	_, err := r.tx.Exec(query,
		item.SellerId,
		item.StartPrice,
		item.CurrentBid,
		item.StartTime,
		item.EndTime,
		item.ExtraTimeEnabled,
		item.ExtraTimeDuration,
		item.ExtraTimeThreshold,
		item.Status,
	)

	return err
}
