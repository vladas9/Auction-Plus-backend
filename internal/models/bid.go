package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type BidModel struct {
	BaseModel
	AuctionId uuid.UUID       `json:"auction_id"`
	UserId    uuid.UUID       `json:"user_id"`
	Amount    decimal.Decimal `json:"amount"`
	Timestamp time.Time       `json:"timestamp"`
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
	host, port, title, category string,
	opened bool,
	endTime time.Time) *BidsTable {

	return &BidsTable{
		ID:       uuid.New(),
		ImgSrc:   fmt.Sprintf("http://%s:%s/api/img/%s", host, port, image.String()),
		LotTitle: title,
		MaxBid:   maxBid,
		EndDate:  endTime,
		Category: category,
		Opened:   opened,
		UsersBid: userBid,
	}

}
