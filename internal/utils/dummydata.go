package server

import (
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	"time"
)

func generateDummyAuctions() []m.AuctionModel {
	auctions := []m.AuctionModel{
		{
			AuctionId:          uuid.New(),
			SellerId:           uuid.New(),
			StartingBid:        m.Dollars{Exact: 100, Cents: 0},
			ClosingBid:         m.Dollars{Exact: 200, Cents: 0},
			StartTime:          time.Now().Add(-24 * time.Hour),
			EndTime:            time.Now().Add(24 * time.Hour),
			ExtraTimeDuration:  5 * time.Minute,
			ExtraTimeThreshold: 10 * time.Minute,
			ExtraTimeEnabled:   true,
			IsActive:           true,
		},
		{
			AuctionId:          uuid.New(),
			SellerId:           uuid.New(),
			StartingBid:        m.Dollars{Exact: 50, Cents: 0},
			ClosingBid:         m.Dollars{Exact: 150, Cents: 0},
			StartTime:          time.Now().Add(-48 * time.Hour),
			EndTime:            time.Now().Add(48 * time.Hour),
			ExtraTimeDuration:  10 * time.Minute,
			ExtraTimeThreshold: 15 * time.Minute,
			ExtraTimeEnabled:   false,
			IsActive:           false,
		},
		{
			AuctionId:          uuid.New(),
			SellerId:           uuid.New(),
			StartingBid:        m.Dollars{Exact: 300, Cents: 0},
			ClosingBid:         m.Dollars{Exact: 350, Cents: 0},
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

func generateDummyUser() m.UserModel {
	return m.UserModel{
		UserId:         uuid.New(),
		Username:       "john_doe",
		Email:          "john.doe@example.com",
		Password:       "hashed_password", // In real case, this would be hashed.
		Address:        "1234 Elm St, Springfield, USA",
		PhoneNumber:    "+1-555-1234",
		UserType:       "seller",                     // can be 'buyer', 'seller', or 'admin'
		RegisteredDate: time.Now().AddDate(0, -6, 0), // registered 6 months ago
	}
}
