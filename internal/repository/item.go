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

func (stx *StoreTx) ItemRepo() *itemRepo {
	return &itemRepo{tx: stx.Tx}
}

func (r *itemRepo) GetById(id uuid.UUID) (*m.ItemModel, error) {
	item := &m.ItemModel{}
	query := `
		SELECT 
			id,
			name,
			description,
			category,
			condition,
			images
		FROM
			Items
		WHERE
			id = $1
	`
	row := r.tx.QueryRow(query, id)
	if err := row.Scan(
		&item.ID,
		&item.Name,
		&item.Description,
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
			id,
			name,
			description,
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
			&item.ID,
			&item.Name,
			&item.Description,
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
			category = $3,
			condition = $4,
			images = $5
		WHERE
			id = $6
	`
	_, err := r.tx.Exec(query,
		item.Name,
		item.Description,
		item.Category,
		item.Condition,
		pq.Array(item.Images),
		item.Id(),
	)

	return err
}

// Remove deletes an item by its ID
func (r *itemRepo) Remove(id uuid.UUID) error {
	query := `
		DELETE FROM 
			Items
		WHERE 
			id = $1
	`
	_, err := r.tx.Exec(query, id)

	return err
}

// Insert adds a new item to the database
func (r *itemRepo) Insert(item *m.ItemModel) (uuid.UUID, error) {
	query := `
		INSERT INTO Items (
			name,
			description,
			category,
			condition,
			images
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id
	`
	var itemId string
	err := r.tx.QueryRow(query,
		item.Name,
		item.Description,
		item.Category,
		item.Condition,
		pq.Array(item.Images),
	).Scan(&itemId)

	if err != nil {
		return uuid.Nil, err
	}

	return uuid.MustParse(itemId), err
}
