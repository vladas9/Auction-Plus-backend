package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type transactionRepo struct {
	tx *sql.Tx
}

// Struct to hold property-value pairs for filtering
type FilterCondition struct {
	Property string
	Value    interface{}
}

func (s *StoreTx) TransactionRepo() *transactionRepo {
	return &transactionRepo{tx: s.Tx}
}

func (r *transactionRepo) GetById(id uuid.UUID) (*m.TransactionModel, error) {
	item := &m.TransactionModel{}
	query := `
		SELECT 
			id,
			auction_id,
			buyer_id,
			seller_id,
			amount,
			date
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
		&item.Date,
	); err != nil {
		return nil, err
	}
	return item, nil
}

// GetAll retrieves transactions based on a set of filter conditions (field and value pairs).
func (r *transactionRepo) GetAll(filters []FilterCondition) ([]*m.TransactionModel, error) {
	var transactions []*m.TransactionModel
	query := `
		SELECT 
			id,
			auction_id,
			buyer_id,
			seller_id,
			amount,
			date
		FROM
			transactions
		WHERE 1=1
	`

	var args []interface{}
	var conditions []string

	// Dynamically build the WHERE clause based on the provided filters
	for i, filter := range filters {
		conditions = append(conditions, fmt.Sprintf("%s = $%d", filter.Property, i+1))
		args = append(args, filter.Value)
	}

	// Append the conditions to the query if any are provided
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.tx.Query(query, args...)
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
			&item.Date,
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

func (r *transactionRepo) Update(item *m.TransactionModel) error {
	query := `
		UPDATE 
			transactions
		SET
			auction_id = $1,
			buyer_id = $2,
			seller_id = $3,
			amount = $4,
			date = $5
		WHERE
			id = $6
	`
	_, err := r.tx.Exec(query,
		item.AuctionId,
		item.BuyerId,
		item.SellerId,
		item.Amount,
		item.Date,
		item.ID,
	)

	return err
}

func (r *transactionRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			transactions
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

func (r *transactionRepo) Insert(item *m.TransactionModel) (uuid.UUID, error) {
	query := `
		INSERT INTO transactions (
			auction_id,
			buyer_id,
			seller_id,
			amount,
			date
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id
	`
	var itemId string
	err := r.tx.QueryRow(query,
		item.AuctionId,
		item.BuyerId,
		item.SellerId,
		item.Amount,
		item.Date,
	).Scan(&itemId)

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(itemId), nil
}
