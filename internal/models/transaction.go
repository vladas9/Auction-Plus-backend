package models

import (
	"github.com/google/uuid"
	"time"
)

type TransactionModel struct {
	BaseModel
	AuctionId       uuid.UUID `json:"auction_id"`
	BuyerId         uuid.UUID `json:"buyer_id"`
	SellerId        uuid.UUID `json:"seller_id"`
	Amount          Decimal   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
}
