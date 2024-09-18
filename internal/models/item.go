package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ItemModel struct {
	BaseModel
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	StartingPrice decimal.Decimal `json:"starting_price"`
	Category      Category        `json:"category"`
	Condition     Condition       `json:"condition"`
	Images        []uuid.UUID     `json:"images"`
}
