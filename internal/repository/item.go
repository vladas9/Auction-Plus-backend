package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lib/pq"
	m "github.com/vladas9/backend-practice/internal/models"
)

type itemRepo struct {
	tx *sql.Tx
}

func ItemRepo(tx *sql.Tx) *itemRepo {
	return &itemRepo{tx}
}

func (r *itemRepo) GetById(id uuid.UUID) (*m.ItemModel, error) {
	item := &m.ItemModel{}
	query := `
		SELECT 
			item_id,
			name,
			description,
			starting_price,
			category,
			condition,
			images
		FROM
			Items
		WHERE
			item_id = $1
	`
	row := r.tx.QueryRow(query, id)
	if err := row.Scan(
		&item.ItemId,
		&item.Name,
		&item.Description,
		&item.StartingPrice,
		&item.Category,
		&item.Condition,
		pq.Array(&item.Images), // pq.Array for handling array in PostgreSQL
	); err != nil {
		return nil, err
	}
	return item, nil
}

// GetAll retrieves all items
func (r *itemRepo) GetAll() ([]*m.ItemModel, error) {
	var items []*m.ItemModel
	query := `
		SELECT 
			item_id,
			name,
			description,
			starting_price,
			category,
			condition,
			images
		FROM
			Items
	`
	rows, err := r.tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := &m.ItemModel{}
		if err := rows.Scan(
			&item.ItemId,
			&item.Name,
			&item.Description,
			&item.StartingPrice,
			&item.Category,
			&item.Condition,
			pq.Array(&item.Images),
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Update modifies an existing item in the database
func (r *itemRepo) Update(item *m.ItemModel) error {
	query := `
		UPDATE 
			Items
		SET
			name = $1,
			description = $2,
			starting_price = $3,
			category = $4,
			condition = $5,
			images = $6
		WHERE
			item_id = $7
	`
	_, err := r.tx.Exec(query,
		item.Name,
		item.Description,
		item.StartingPrice,
		item.Category,
		item.Condition,
		pq.Array(item.Images),
		item.ItemId,
	)

	return err
}

// Remove deletes an item by its ID
func (r *itemRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			Items
		WHERE 
			item_id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

// Insert adds a new item to the database
func (r *itemRepo) Insert(item *m.ItemModel) error {
	query := `
		INSERT INTO Items (
			item_id,
			name,
			description,
			starting_price,
			category,
			condition,
			images
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`
	_, err := r.tx.Exec(query,
		item.ItemId,
		item.Name,
		item.Description,
		item.StartingPrice,
		item.Category,
		item.Condition,
		pq.Array(item.Images),
	)

	return err
}
