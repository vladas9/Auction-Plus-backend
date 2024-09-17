package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	m "github.com/vladas9/backend-practice/internal/models"
)

func GenerateDummyAuctions(sellerId uuid.UUID) []*m.AuctionModel {
	auctions := []*m.AuctionModel{
		{
			BaseModel:          m.BaseModel{uuid.New()},
			SellerId:           sellerId,
			StartPrice:         decimal.NewFromInt(100),
			CurrentBid:         decimal.NewFromInt(200),
			StartTime:          time.Now().Add(-24 * time.Hour),
			EndTime:            time.Now().Add(24 * time.Hour),
			ExtraTimeDuration:  5 * time.Minute,
			ExtraTimeThreshold: 10 * time.Minute,
			ExtraTimeEnabled:   true,
			Status:             true,
		},
		{
			BaseModel:          m.BaseModel{uuid.New()},
			SellerId:           sellerId,
			StartPrice:         decimal.NewFromInt(50),
			CurrentBid:         decimal.NewFromInt(150),
			StartTime:          time.Now().Add(-48 * time.Hour),
			EndTime:            time.Now().Add(48 * time.Hour),
			ExtraTimeDuration:  10 * time.Minute,
			ExtraTimeThreshold: 15 * time.Minute,
			ExtraTimeEnabled:   false,
			Status:             false,
		},
		{
			BaseModel:          m.BaseModel{uuid.New()},
			SellerId:           sellerId,
			StartPrice:         decimal.NewFromInt(300),
			CurrentBid:         decimal.NewFromInt(350),
			StartTime:          time.Now().Add(-72 * time.Hour),
			EndTime:            time.Now().Add(72 * time.Hour),
			ExtraTimeDuration:  3 * time.Minute,
			ExtraTimeThreshold: 5 * time.Minute,
			ExtraTimeEnabled:   true,
			Status:             true,
		},
	}

	return auctions
}

func GenerateDummyUser() *m.UserModel {
	return &m.UserModel{
		Username:       "john_doe",
		Email:          "john.doe@example.com",
		Image:          "",
		Password:       "hashed_password", // In real case, this would be hashed.
		Address:        "1234 Elm St, Springfield, USA",
		PhoneNumber:    "+1-555-1234",
		UserType:       "seller",                     // can be 'buyer', 'seller', or 'admin'
		RegisteredDate: time.Now().AddDate(0, -6, 0), // registered 6 months ago
	}
}

func GenerateDummyBids() []m.BidModel {
	return []m.BidModel{
		m.BidModel{
			BaseModel: m.BaseModel{uuid.New()},
			UserId:    uuid.New(),
			Amount:    decimal.NewFromInt(150),
			Timestamp: time.Now(),
		},
		m.BidModel{
			BaseModel: m.BaseModel{uuid.New()},
			UserId:    uuid.New(),
			Amount:    decimal.NewFromInt(250),
			Timestamp: time.Now(),
		},
		m.BidModel{
			BaseModel: m.BaseModel{uuid.New()},
			UserId:    uuid.New(),
			Amount:    decimal.NewFromInt(300),
			Timestamp: time.Now(),
		},
		m.BidModel{
			BaseModel: m.BaseModel{uuid.New()},
			UserId:    uuid.New(),
			Amount:    decimal.NewFromInt(400),
			Timestamp: time.Now(),
		},
	}
}
