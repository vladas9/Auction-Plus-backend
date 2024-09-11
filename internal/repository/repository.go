package repository

import (
	//"database/sql"
	//"fmt"
	//"strings"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type Repo[T m.Model] interface {
	GetAll() ([]*T, error)
	GetById(id uuid.UUID) (*T, error)
	Update(model *T) error
	Insert(model *T) (any, error)
	Delete(model *T) error
}
