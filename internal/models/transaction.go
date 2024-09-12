package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionModel struct {
	BaseModel
	AuctionId       uuid.UUID       `json:"auction_id"`
	BuyerId         uuid.UUID       `json:"buyer_id"`
	SellerId        uuid.UUID       `json:"seller_id"`
	Amount          decimal.Decimal `json:"amount"`
	TransactionDate time.Time       `json:"transaction_date"`
}
