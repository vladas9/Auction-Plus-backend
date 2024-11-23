package repository

import (
	"database/sql"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type NotificationRepo struct {
	tx *sql.Tx
}

func (r *NotificationRepo) GetAll() ([]*m.NotificationModel, error) { return nil, nil }

func (r *NotificationRepo) GetById(id uuid.UUID) (*m.NotificationModel, error) { return nil, nil }

func (r *NotificationRepo) Update(model *m.NotificationModel) error { return nil }

func (r *NotificationRepo) Insert(model *m.NotificationModel) error { return nil }

func (r *NotificationRepo) Delete(model *m.NotificationModel) error { return nil }
