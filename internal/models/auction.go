package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AuctionModel struct {
	BaseModel
	SellerId           uuid.UUID       `json:"seller_id"`
	CurrentBid         decimal.Decimal `json:"current_bid"`
	StartTime          time.Time       `json:"start_time"`
	EndTime            time.Time       `json:"end_time"`
	ExtraTimeDuration  time.Duration   `json:"extra_time_duration"`
	ExtraTimeThreshold time.Duration   `json:"extra_time_threshold"`
	ExtraTimeEnabled   bool            `json:"extra_time_enabled"`
	Status             bool            `json:"status"`
}
