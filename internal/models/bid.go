package models

import (
	"github.com/google/uuid"
	"time"
)

type Bid struct {
	BidId     uuid.UUID `json:"bid_is"`
	AuctionId uuid.UUID `json:"auction_id"`
	UserId    uuid.UUID `json:"user_id"`
	Amount    Dollars   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
