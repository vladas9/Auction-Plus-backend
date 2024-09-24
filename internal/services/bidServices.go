package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vladas9/backend-practice/internal/dtos"
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
)

func (s *Service) NewBid(bid *m.BidModel) (err error) {
	auction := &m.AuctionModel{}
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		auction, err = stx.AuctionRepo().GetById(bid.AuctionId)
		if err != nil {
			return err
		}

		if auction.CurrentBid.Compare(bid.Amount) == -1 {
			err = stx.BidRepo().Insert(bid)
			if err != nil {
				return fmt.Errorf("failed to insert bid: %w", err)
			}
		} else {
			return fmt.Errorf("Bid is smaller than CurrentBid: %s", bid.Amount)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("Palacing bid failed: %s", err)
	}

	return nil
}

func (s *Service) GetBidTable(userId uuid.UUID, limit, offset int) ([]*dtos.BidsTable, error) {
	var bidList []*m.BidModel
	var auctionList []*m.AuctionModel
	var itemList []*m.ItemModel
	var userList []*m.UserModel

	err := s.store.WithTx(func(stx *r.StoreTx) error {
		var err error
		bidList, err = stx.BidRepo().GetAllByUserId(userId, limit, offset)
		if err != nil {
			return fmt.Errorf("Failed getting bids: %s", err)
		}

		auctionList, itemList, userList, err = getRelatedData(stx, bidList)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	bidsTable, err := buildBidsTable(bidList, auctionList, itemList, userList)
	if err != nil {
		return nil, err
	}

	return bidsTable, nil
}

func getRelatedData(stx *r.StoreTx, bidList []*m.BidModel) (
	[]*m.AuctionModel, []*m.ItemModel, []*m.UserModel, error,
) {
	auctionList := make([]*m.AuctionModel, 0)
	itemList := make([]*m.ItemModel, 0)
	userList := make([]*m.UserModel, 0)

	auctionMap := make(map[uuid.UUID]*m.AuctionModel)
	for _, bid := range bidList {
		auctionId := bid.AuctionId
		auction, err := stx.AuctionRepo().GetById(auctionId)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("Failed getting auction: %s", err)
		}
		auctionList = append(auctionList, auction)
		auctionMap[auctionId] = auction
	}

	for _, auction := range auctionList {
		itemId := auction.ItemId
		item, err := stx.ItemRepo().GetById(itemId)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("Failed getting item: %s", err)
		}
		itemList = append(itemList, item)

		userId := auction.MaxBidderId
		user, err := stx.UserRepo().GetByProperty("id", userId)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("Failed getting user: %s", err)
		}
		userList = append(userList, user)
	}

	return auctionList, itemList, userList, nil
}

func buildBidsTable(
	bidList []*m.BidModel,
	auctionList []*m.AuctionModel,
	itemList []*m.ItemModel,
	userList []*m.UserModel,
) ([]*dtos.BidsTable, error) {
	userMap := CreateUserMap(userList)
	auctionMap := CreateAuctionMap(auctionList)
	itemMap := CreateItemMap(itemList)
	highestBids := FindHighestBids(bidList)

	bidsTable := make([]*dtos.BidsTable, 0)
	for _, bid := range highestBids {
		relatedAuction, auctionExists := auctionMap[bid.AuctionId]
		if !auctionExists {
			continue
		}

		relatedItem, itemExists := itemMap[relatedAuction.ItemId]
		if !itemExists {
			continue
		}

		relatedUser, userExists := userMap[relatedAuction.MaxBidderId]
		if !userExists {
			continue
		}

		image := m.GetFirstImageOrNil(relatedItem)

		bidTableEntry := dtos.BidsTableMapper(
			image,
			relatedAuction.CurrentBid,
			bid.Amount,
			Host,
			Port,
			relatedItem.Name,
			string(relatedItem.Category),
			relatedUser.Username,
			relatedAuction.Status,
			relatedAuction.EndTime,
		)

		bidsTable = append(bidsTable, bidTableEntry)
	}

	return bidsTable, nil
}
