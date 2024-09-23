package services

import (
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"

	"fmt"
	"runtime"
)

func fail(err error) error {
	if err == nil {
		return nil
	}

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("unknown function: %w", err)
	}
	fn := runtime.FuncForPC(pc)
	functionName := fn.Name()

	return fmt.Errorf("%s: %w", functionName, err)
}

func CreateUserMap(userList []*m.UserModel) map[uuid.UUID]*m.UserModel {
	userMap := make(map[uuid.UUID]*m.UserModel)
	for _, user := range userList {
		userMap[user.ID] = user
	}
	return userMap
}

func CreateAuctionMap(auctionList []*m.AuctionModel) map[uuid.UUID]*m.AuctionModel {
	auctionMap := make(map[uuid.UUID]*m.AuctionModel)
	for _, auction := range auctionList {
		auctionMap[auction.ID] = auction
	}
	return auctionMap
}

func CreateItemMap(itemList []*m.ItemModel) map[uuid.UUID]*m.ItemModel {
	itemMap := make(map[uuid.UUID]*m.ItemModel)
	for _, item := range itemList {
		itemMap[item.ID] = item
	}
	return itemMap
}

func FindHighestBids(bidList []*m.BidModel) map[uuid.UUID]*m.BidModel {
	highestBids := make(map[uuid.UUID]*m.BidModel)
	for _, bid := range bidList {
		if existingBid, exists := highestBids[bid.AuctionId]; !exists || bid.Amount.Compare(existingBid.Amount) == 1 {
			highestBids[bid.AuctionId] = bid
		}
	}
	return highestBids
}
