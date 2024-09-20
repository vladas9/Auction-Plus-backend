package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type AuctionModel struct {
	BaseModel
	SellerId           uuid.UUID       `json:"seller_id"`
	ItemId             uuid.UUID       `json:"item_id"`
	StartPrice         decimal.Decimal `json:"starting_price"`
	CurrentBid         decimal.Decimal `json:"current_bid"`
	MaxBidderId        uuid.UUID       `json:"max_bidder_id"`
	BidCount           int16           `json:"bid_count"`
	StartTime          time.Time       `json:"start_time"`
	EndTime            time.Time       `json:"end_time"`
	ExtraTimeDuration  time.Duration   `json:"extra_time_duration"`
	ExtraTimeThreshold time.Duration   `json:"extra_time_threshold"`
	ExtraTimeEnabled   bool            `json:"extra_time_enabled"`
	Status             bool            `json:"status"`
}

type AuctionTable struct {
	ID        uuid.UUID       `json:"id"`
	ImgSrc    string          `json:"img_src"`
	LotTitle  string          `json:"lot_title"`
	MaxBid    decimal.Decimal `json:"max_bid"`
	EndDate   time.Time       `json:"end_date"`
	Category  string          `json:"category"`
	Opened    bool            `json:"opened"`
	TopBidder string          `json:"top_bidder"`
}

func AuctionTableMapper(
	image uuid.UUID,
	maxBid decimal.Decimal,
	host, port, title, category, username string,
	opened bool,
	endTime time.Time) *AuctionTable {

	return &AuctionTable{
		ID:        uuid.New(),
		ImgSrc:    fmt.Sprintf("http://%s:%s/api/img/%s", host, port, image.String()),
		LotTitle:  title,
		MaxBid:    maxBid,
		EndDate:   endTime,
		Category:  category,
		Opened:    opened,
		TopBidder: username,
	}

}
