package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	m "github.com/vladas9/backend-practice/internal/models"
	"github.com/vladas9/backend-practice/internal/utils"
)

var (
	Images = []uuid.UUID{uuid.New(), uuid.New(), uuid.New(), uuid.New()}
)

func rantBool() bool {
	if rand.Float32() < 0.5 {
		return false
	}
	return true
}

func randCond() m.Condition {
	if rantBool() {
		return m.New
	}
	return m.Used
}

func randCateg() m.Category {
	categories := []m.Category{m.Electronics, m.Furniture, m.Arts, m.RealEstate, m.Other}
	return categories[rand.Intn(4)]
}

func CreateDummyItem() *m.ItemModel {
	itemUUID := uuid.New()

	dummyItem := &m.ItemModel{
		BaseModel:   m.BaseModel{ID: itemUUID},
		Name:        "Dummy Item",
		Description: "This is a dummy item.",
		Category:    m.Other,
		Condition:   randCond(),
		Images:      Images,
	}
	return dummyItem
}

func CreateDummyAuction(itemId, userId uuid.UUID) *m.AuctionModel {
	startPrice := 10 + int64(rand.Intn(1000))
	return &m.AuctionModel{
		BaseModel:          m.BaseModel{},
		SellerId:           userId,
		ItemId:             itemId,
		StartPrice:         decimal.NewFromInt(startPrice),
		CurrentBid:         decimal.Zero,
		BidCount:           0,
		StartTime:          time.Now().Add(-72 * time.Hour),
		EndTime:            time.Now().Add(72 * time.Hour),
		ExtraTimeDuration:  3 * time.Minute,
		ExtraTimeThreshold: 5 * time.Minute,
		ExtraTimeEnabled:   true,
		Status:             true,
	}
}

func CreateDummyTransaction(auctionId, sellerId, buyerId uuid.UUID) *m.TransactionModel {
	return &m.TransactionModel{
		BaseModel: m.BaseModel{},
		AuctionId: auctionId,
		BuyerId:   buyerId,
		SellerId:  sellerId,
		Amount:    decimal.NewFromInt(400),
		Date:      time.Now(),
	}
}

func GenerateDummyUser(email string) *m.UserModel {
	user := &m.UserModel{
		BaseModel:      m.BaseModel{},
		Username:       "john_doe",
		Email:          email,
		Image:          uuid.NewString(),
		Password:       "password",
		Address:        "1234 Elm St, Springfield, USA",
		PhoneNumber:    "+1-555-1234",
		UserType:       "admin",                      // admin or client
		RegisteredDate: time.Now().AddDate(0, -6, 0), // registered 6 months ago
	}
	user.Password, _ = utils.HashPassword(user.Password)
	return user
}

func GenerateDummyBids(bidder, auct uuid.UUID, price decimal.Decimal) []*m.BidModel {
	return []*m.BidModel{
		&m.BidModel{
			BaseModel: m.BaseModel{},
			AuctionId: auct,
			UserId:    bidder,
			Amount:    price.Add(decimal.NewFromInt(150)),
			Timestamp: time.Now(),
		},
		&m.BidModel{
			BaseModel: m.BaseModel{},
			AuctionId: auct,
			UserId:    bidder,
			Amount:    price.Add(decimal.NewFromInt(250)),
			Timestamp: time.Now(),
		},
		&m.BidModel{
			BaseModel: m.BaseModel{},
			AuctionId: auct,
			UserId:    bidder,
			Amount:    price.Add(decimal.NewFromInt(300)),
			Timestamp: time.Now(),
		},
		&m.BidModel{
			BaseModel: m.BaseModel{},
			AuctionId: auct,
			UserId:    bidder,
			Amount:    price.Add(decimal.NewFromInt(400)),
			Timestamp: time.Now(),
		},
	}
}
