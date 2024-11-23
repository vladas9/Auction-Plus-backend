package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	u "github.com/vladas9/backend-practice/internal/utils"
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

type AuctionDetails struct {
	Auction   *AuctionModel
	Item      *ItemModel
	BidList   []*BidModel
	MaxBidder *UserModel
}

func NewAuctionDetails(auct *AuctionModel) *AuctionDetails {
	return &AuctionDetails{
		Auction:   auct,
		Item:      nil,
		BidList:   nil,
		MaxBidder: nil,
	}
}

func (rsp *AuctionDetails) ItemHas(condition, category string) bool {
	hasCateg := (category == "" || rsp.Item.Category == Category(category))
	hasCond := (condition == "" || rsp.Item.Condition == Condition(condition))

	u.Logger.Info("Checking item against filters:",
		"conditionMet:", hasCond,
		"condition:", condition,
		"categoryMet:", hasCateg,
		"category:", category)

	return hasCateg && hasCond
}
