package models

import (
	"github.com/google/uuid"
	"time"
)

type BidModel struct {
	BaseModel
	AuctionId uuid.UUID `json:"auction_id"`
	UserId    uuid.UUID `json:"user_id"`
	Amount    Decimal   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
