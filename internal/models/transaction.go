package models

import (
	"github.com/google/uuid"
	"time"
)

type TransactionModel struct {
	TransactionId   uuid.UUID `json:"transaction_id"`
	AuctionId       uuid.UUID `json:"auction_id"`
	BuyerId         uuid.UUID `json:"buyer_id"`
	SellerId        uuid.UUID `json:"seller_id"`
	Amount          int64     `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
}
