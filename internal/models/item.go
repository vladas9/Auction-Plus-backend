package models

import (
	"github.com/google/uuid"
)

type ItemModel struct {
	ItemId        uuid.UUID `json:"item_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	StartingPrice Decimal   `json:"starting_price"`
	Category      Category  `json:"category"`
	Condition     Condition `json:"condition"`
	Images        []string  `json:"images"`
}
