package dtos

import (
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"

	"fmt"
	"os"
	"time"
)

//{
//    "id": 1,
//    "img_srcs": [
//      "https://.jpg",
//      "https://.jpg",
//      "https://.jpg",
//      "https://encry"
//    ],
//    "title": "Test lot",
//    "max_bid": 300,
//    "end_date": "2024-09-09T13:34:15+03:00",
//    "n_bids": 20,
//    "opened": true,
//    "description": "loremohegoreoubueoiwbbvnre",
//    "labels": ["11 January", "12 January", "13 January"],
//    "bids_perday": [12, 23, 43],
//    "max_bid_perday": [23, 45, 700]
//
//}

type AuctionFull struct {
	Id          uuid.UUID `json:"id"`
	ImgSrc      []string  `json:"img_src"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Opened      bool      `json:"opened"`

	Category  m.Category  `json:"category_name"`
	Condition m.Condition `json:"condition"`

	EndDate   time.Time `json:"end_date"`
	NumOfBids int       `json:"n_bids"`
	MaxBid    m.Decimal `json:"max_bid"`

	Labels     []string    `json:"labels"`
	BidsPerDay []int       `json:"bids_perday"`
	MaxBids    []m.Decimal `json:"max_bid_perday"`
}

func MapAuctionRespToFull(auctDets *m.AuctionDetails) *AuctionFull {
	item := auctDets.Item
	auction := auctDets.Auction

	bidStats := m.GetBidStats(auctDets.BidList)

	data := &AuctionFull{
		Id:          auction.Id(),
		ImgSrc:      MakeImgSrcs(item.Images),
		Title:       item.Name,
		Description: item.Description,
		Opened:      auction.Status,
		Category:    item.Category,
		Condition:   item.Condition,
		EndDate:     auction.EndTime,
		MaxBid:      auction.CurrentBid,
		NumOfBids:   int(auction.BidCount),
		Labels:      bidStats.Labels,
		BidsPerDay:  bidStats.BidsCount,
		MaxBids:     bidStats.MaxBids,
	}
	return data
}
func MakeImgSrcs(ids []uuid.UUID) []string {
	var srcList []string
	for _, id := range ids {
		src := fmt.Sprintf("http://%s:%s/api/img/%s", os.Getenv("HOST"), os.Getenv("PORT"), id)
		srcList = append(srcList, src)
	}
	return srcList
}
