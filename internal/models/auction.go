package models

import (
	"github.com/google/uuid"
	"time"
)

type Dollars struct {
	Exact, Cents uint32
}

type AuctionModel struct {
	AuctionId          uuid.UUID     `json:"auction_id"`
	SellerId           uuid.UUID     `json:"seller_id"`
	StartingBid        Dollars       `json:"starting_bid"`
	ClosingBid         Dollars       `json:"closing_bid"`
	StartTime          time.Time     `json:"start_time"`
	EndTime            time.Time     `json:"end_time"`
	ExtraTimeDuration  time.Duration `json:"extra_time_duration"`
	ExtraTimeThreshold time.Duration `json:"extra_time_threshold"`
	ExtraTimeEnabled   bool          `json:"extra_time_enabled"`
	IsActive           bool          `json:"is_active"`
}
