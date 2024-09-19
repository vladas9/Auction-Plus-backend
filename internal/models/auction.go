package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AuctionModel struct {
	BaseModel
	SellerId           uuid.UUID       `json:"seller_id"`
	ItemId             uuid.UUID       `json:"item_id"`
	StartingBid        decimal.Decimal `json:"starting_bid"`
	CurrentBid         decimal.Decimal `json:"current_bid"`
	MaxBidderId        uuid.UUID       `json:"max_bidder_id"`
	BidCount           int16           `json:"bid_count"`
	StartTime          time.Time       `json:"start_time"`
	EndTime            time.Time       `json:"end_time"`
	ExtraTimeDuration  time.Duration   `json:"extra_time_duration"`
	ExtraTimeThreshold time.Duration   `json:"extra_time_threshold"`
	ExtraTimeEnabled   bool            `json:"extra_time_enabled"`
	IsActive           bool            `json:"is_active"`
}
