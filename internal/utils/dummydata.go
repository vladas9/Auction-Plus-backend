package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	m "github.com/vladas9/backend-practice/internal/models"
)

func GenerateDummyAuctions() []m.AuctionModel {
	auctions := []m.AuctionModel{
		{
			BaseModel:          m.BaseModel{uuid.New()},
			SellerId:           uuid.New(),
			StartingBid:        decimal.NewFromInt(100),
			ClosingBid:         decimal.NewFromInt(200),
			StartTime:          time.Now().Add(-24 * time.Hour),
			EndTime:            time.Now().Add(24 * time.Hour),
			ExtraTimeDuration:  5 * time.Minute,
			ExtraTimeThreshold: 10 * time.Minute,
			ExtraTimeEnabled:   true,
			IsActive:           true,
		},
		{
			BaseModel:          m.BaseModel{uuid.New()},
			SellerId:           uuid.New(),
			StartingBid:        decimal.NewFromInt(50),
			ClosingBid:         decimal.NewFromInt(150),
			StartTime:          time.Now().Add(-48 * time.Hour),
			EndTime:            time.Now().Add(48 * time.Hour),
			ExtraTimeDuration:  10 * time.Minute,
			ExtraTimeThreshold: 15 * time.Minute,
			ExtraTimeEnabled:   false,
			IsActive:           false,
		},
		{
			BaseModel:          m.BaseModel{uuid.New()},
			SellerId:           uuid.New(),
			StartingBid:        decimal.NewFromInt(300),
			ClosingBid:         decimal.NewFromInt(350),
			StartTime:          time.Now().Add(-72 * time.Hour),
			EndTime:            time.Now().Add(72 * time.Hour),
			ExtraTimeDuration:  3 * time.Minute,
			ExtraTimeThreshold: 5 * time.Minute,
			ExtraTimeEnabled:   true,
			IsActive:           true,
		},
	}

	return auctions
}

func GenerateDummyUser() *m.UserModel {
	return &m.UserModel{
		BaseModel:      m.BaseModel{uuid.New()},
		Username:       "john_doe",
		Email:          "john.doe@example.com",
		Password:       "hashed_password", // In real case, this would be hashed.
		Address:        "1234 Elm St, Springfield, USA",
		PhoneNumber:    "+1-555-1234",
		UserType:       "seller",                     // can be 'buyer', 'seller', or 'admin'
		RegisteredDate: time.Now().AddDate(0, -6, 0), // registered 6 months ago
	}
}
