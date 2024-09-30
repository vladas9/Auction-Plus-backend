package repository

import (
	"database/sql"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	"github.com/vladas9/backend-practice/internal/utils"
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
			item_id,
			start_price,
			current_bid,
			max_bidder_id,
			bid_count,
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
		&item.ItemId,
		&item.StartPrice,
		&item.CurrentBid,
		&item.MaxBidderId,
		&item.BidCount,
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

func (r *auctionRepo) GetAllByUserId(id uuid.UUID, limit, offset int) ([]*m.AuctionModel, error) {
	var auctions []*m.AuctionModel
	query := `
		SELECT 
			id,
			seller_id,
			item_id,
			start_price,
			current_bid,
			max_bidder_id,
			bid_count,
			start_time,
			end_time,
			extra_time_enabled,
			extra_time_duration,
			extra_time_threshold,
			status
		FROM
			auctions
		WHERE
			seller_id = $1
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
		item := &m.AuctionModel{}
		if err := rows.Scan(
			&item.ID,
			&item.SellerId,
			&item.ItemId,
			&item.StartPrice,
			&item.CurrentBid,
			&item.MaxBidderId,
			&item.BidCount,
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

func (r *auctionRepo) GetAllFiltered(offset, limit int, minPrice, maxPrice m.Decimal) ([]*m.AuctionModel, error) {
	var auctions []*m.AuctionModel
	query := `
		SELECT 
			id,
			seller_id,
			item_id,
			start_price,
			current_bid,
			bid_count,
			start_time,
			end_time,
			extra_time_enabled,
			extra_time_duration,
			extra_time_threshold,
			status
		FROM
			auctions
		WHERE
			($1 <= current_bid) AND ($2 = 0 OR $2 >= current_bid)
		LIMIT
			CASE WHEN $3 = 0 THEN 100 ELSE $3 END
		OFFSET 
			$4
	`
	utils.Logger.Info(minPrice, maxPrice)
	rows, err := r.tx.Query(query, minPrice, maxPrice, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		utils.Logger.Info("rows next")
		item := &m.AuctionModel{}
		if err := rows.Scan(
			&item.ID,
			&item.SellerId,
			&item.ItemId,
			&item.StartPrice,
			&item.CurrentBid,
			&item.BidCount,
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
	utils.Logger.Info("AuctionRepo GetAllFiltered:", auctions)
	return auctions, nil
}

func (r *auctionRepo) GetAll(offset, limit int) ([]*m.AuctionModel, error) {
	var auctions []*m.AuctionModel
	query := `
		SELECT 
			id,
			seller_id,
			item_id,
			start_price,
			current_bid,
			bid_count,
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
	rows, err := r.tx.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		item := &m.AuctionModel{}
		if err := rows.Scan(
			&item.ID,
			&item.SellerId,
			&item.ItemId,
			&item.StartPrice,
			&item.CurrentBid,
			&item.BidCount,
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
			bid_count = $4,
			start_time = $5,
			end_time = $6,
			extra_time_enabled = $7,
			extra_time_duration = $8,
			extra_time_threshold = $9,
			status = $10
			item_id = $11
			max_bidder_id = $12
		WHERE
			id = $13
	`
	_, err := r.tx.Exec(query,
		item.SellerId,
		item.StartPrice,
		item.CurrentBid,
		item.BidCount,
		item.StartTime,
		item.EndTime,
		item.ExtraTimeEnabled,
		item.ExtraTimeDuration,
		item.ExtraTimeThreshold,
		item.Status,
		item.ItemId,
		item.MaxBidderId,
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

func (r *auctionRepo) Insert(item *m.AuctionModel) (uuid.UUID, error) {
	query := `
		INSERT INTO auctions (
			seller_id,
			item_id,
			start_price,
			current_bid,
			max_bidder_id,
			bid_count,
			start_time,
			end_time,
			extra_time_enabled,
			extra_time_duration,
			extra_time_threshold,
			status
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		) RETURNING id
	`
	var itemId string
	err := r.tx.QueryRow(query,
		item.SellerId,
		item.ItemId,
		item.StartPrice,
		item.CurrentBid,
		item.MaxBidderId,
		item.BidCount,
		item.StartTime,
		item.EndTime,
		item.ExtraTimeEnabled,
		item.ExtraTimeDuration,
		item.ExtraTimeThreshold,
		item.Status,
	).Scan(&itemId)

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(itemId), nil
}
