package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type BidModel struct {
	BaseModel
	AuctionId uuid.UUID `json:"auction_id"`
	UserId    uuid.UUID `json:"user_id"`
	Amount    Decimal   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type BidStats struct {
	Labels    []string
	BidsCount []int
	MaxBids   []Decimal
}

func GetBidStats(bids []*BidModel) BidStats {
	stats := BidStats{}

	bidsPerDay := make(map[string][]Decimal)

	for _, bid := range bids {
		date := bid.Timestamp.Format("2 January")
		bidsPerDay[date] = append(bidsPerDay[date], bid.Amount)
	}

	for date, amounts := range bidsPerDay {
		stats.Labels = append(stats.Labels, date)
		stats.BidsCount = append(stats.BidsCount, len(amounts))

		maxAmount := decimal.Zero
		for _, amount := range amounts {
			if amount.GreaterThan(maxAmount) {
				maxAmount = amount
			}
		}
		stats.MaxBids = append(stats.MaxBids, maxAmount)
	}

	return stats
}
