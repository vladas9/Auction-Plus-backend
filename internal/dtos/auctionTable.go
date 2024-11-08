package dtos

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
)

type AuctionTable struct {
	ID        uuid.UUID `json:"id"`
	ImgSrc    string    `json:"img_src"`
	LotTitle  string    `json:"lot_title"`
	MaxBid    m.Decimal `json:"max_bid"`
	EndDate   time.Time `json:"end_date"`
	Category  string    `json:"category"`
	Opened    bool      `json:"opened"`
	TopBidder string    `json:"top_bidder"`
}

func MapAuctionTable(auct *m.AuctionDetails) *AuctionTable {

	return &AuctionTable{
		ID: uuid.New(),
		ImgSrc: fmt.Sprintf("http://%s:%s/api/img/%s",
			os.Getenv("HOST"), os.Getenv("PORT"), m.GetFirstImageOrNil(auct.Item).String()),
		LotTitle:  auct.Item.Name,
		MaxBid:    auct.Auction.CurrentBid,
		EndDate:   auct.Auction.EndTime,
		Category:  string(auct.Item.Category),
		Opened:    auct.Auction.Status,
		TopBidder: auct.MaxBidder.Username,
	}

}
