package models

import (
	"github.com/google/uuid"
	"time"
)

type AuctionModel struct {
	BaseModel
	SellerId           uuid.UUID     `json:"seller_id"`
	StartingBid        Decimal       `json:"starting_bid"`
	ClosingBid         Decimal       `json:"closing_bid"`
	StartTime          time.Time     `json:"start_time"`
	EndTime            time.Time     `json:"end_time"`
	ExtraTimeDuration  time.Duration `json:"extra_time_duration"`
	ExtraTimeThreshold time.Duration `json:"extra_time_threshold"`
	ExtraTimeEnabled   bool          `json:"extra_time_enabled"`
	IsActive           bool          `json:"is_active"`
}
