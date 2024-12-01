package repository

import (
	"database/sql"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type ShippingRepo struct {
	tx *sql.Tx
}

func (r *ShippingRepo) GetAll() ([]*m.ShippingModel, error) { return nil, nil }

func (r *ShippingRepo) GetById(id uuid.UUID) (*m.ShippingModel, error) { return nil, nil }

func (r *ShippingRepo) Update(model *m.ShippingModel) error { return nil }

func (r *ShippingRepo) Insert(model *m.ShippingModel) error { return nil }

func (r *ShippingRepo) Delete(model *m.ShippingModel) error { return nil }
