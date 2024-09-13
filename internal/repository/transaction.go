package repository

import (
	"database/sql"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type TransactionRepo struct {
	tx *sql.Tx
}

func NewTransactionRepo(tx *sql.Tx) *TransactionRepo {
	return &TransactionRepo{tx}
}

func (r *TransactionRepo) GetById(id uuid.UUID) (*m.TransactionModel, error) {
	item := &m.TransactionModel{}
	query := `
		SELECT 
			id,
			auction_id,
			buyer_id,
			seller_id,
			amount,
			transaction_date
		FROM
			transactions
		WHERE
			id = $1
	`
	row := r.tx.QueryRow(query, id)
	if err := row.Scan(
		&item.ID,
		&item.AuctionId,
		&item.BuyerId,
		&item.SellerId,
		&item.Amount,
		&item.TransactionDate,
	); err != nil {
		return nil, err
	}
	return item, nil
}

func (r *TransactionRepo) GetAll() ([]*m.TransactionModel, error) {
	var transactions []*m.TransactionModel
	query := `
		SELECT 
			id,
			auction_id,
			buyer_id,
			seller_id,
			amount,
			transaction_date
		FROM
			transactions
	`
	rows, err := r.tx.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		item := &m.TransactionModel{}
		if err := rows.Scan(
			&item.ID,
			&item.AuctionId,
			&item.BuyerId,
			&item.SellerId,
			&item.Amount,
			&item.TransactionDate,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepo) Update(item *m.TransactionModel) error {
	query := `
		UPDATE 
			transactions
		SET
			auction_id = $1,
			buyer_id = $2,
			seller_id = $3,
			amount = $4,
			transaction_date = $5
		WHERE
			id = $6
	`
	_, err := r.tx.Exec(query,
		&item.AuctionId,
		&item.BuyerId,
		&item.SellerId,
		&item.Amount,
		&item.TransactionDate,
		&item.ID,
	)

	return err
}

func (r *TransactionRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			transactions
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

func (r *TransactionRepo) Insert(item *m.TransactionModel) error {
	query := `
		INSERT INTO transactions (
			id,
			auction_id,
			buyer_id,
			seller_id,
			amount,
			transaction_date
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := r.tx.Exec(query,
		&item.ID,
		&item.AuctionId,
		&item.BuyerId,
		&item.SellerId,
		&item.Amount,
		&item.TransactionDate,
	)

	return err
}
