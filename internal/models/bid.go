package models

import (
	"fmt"
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

type BidsTable struct {
	ID        uuid.UUID       `json:"id"`
	ImgSrc    string          `json:"img_src"`
	LotTitle  string          `json:"lot_title"`
	MaxBid    decimal.Decimal `json:"max_bid"`
	EndDate   time.Time       `json:"end_date"`
	Category  string          `json:"category"`
	Opened    bool            `json:"opened"`
	TopBidder string          `json:"top_bidder"`
	UsersBid  decimal.Decimal `json:"users_bid"`
}

func BidsTableMapper(
	image uuid.UUID,
	maxBid, userBid decimal.Decimal,
	host, port, title, category, username string,
	opened bool,
	endTime time.Time) *BidsTable {

	return &BidsTable{
		ID:        uuid.New(),
		ImgSrc:    fmt.Sprintf("http://%s:%s/api/img/%s", host, port, image.String()),
		LotTitle:  title,
		MaxBid:    maxBid,
		EndDate:   endTime,
		Category:  category,
		Opened:    opened,
		TopBidder: username,
		UsersBid:  userBid,
	}

}
