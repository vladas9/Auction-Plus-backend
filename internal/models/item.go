package models

import (
	"github.com/google/uuid"
)

type ItemModel struct {
	BaseModel
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Category    Category    `json:"category"`
	Condition   Condition   `json:"condition"`
	Images      []uuid.UUID `json:"images"`
}
