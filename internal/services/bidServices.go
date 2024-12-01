package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vladas9/backend-practice/internal/dtos"
	"github.com/vladas9/backend-practice/internal/errors"
	m "github.com/vladas9/backend-practice/internal/models"
	repo "github.com/vladas9/backend-practice/internal/repository"
)

func NewBid(bid *m.BidModel) (err error) {
	auction := &m.AuctionModel{}
	err = repo.WithTx(func(stx *repo.StoreTx) error {
		auction, err = stx.AuctionRepo().GetById(bid.AuctionId)
		if err != nil {
			return errors.NotValid("auction id not valid", err)
		}

		if auction.CurrentBid.Compare(bid.Amount) == -1 {
			err = stx.BidRepo().Insert(bid)
			if err != nil {
				return errors.Internal(err)
			}
		} else {
			return errors.Conflict("Attempt to place bid smaller than current bid", err)
		}
		return nil
	})
	return err
}

func GetBidTable(userId uuid.UUID, limit, offset int) ([]*dtos.BidsTable, error) {
	var bidList []*m.BidModel
	var auctionList []*m.AuctionModel
	var itemList []*m.ItemModel
	var userList []*m.UserModel

	err := repo.WithTx(func(stx *repo.StoreTx) error {
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
		return nil, errors.Internal(err)
	}

	return buildBidsTable(bidList, auctionList, itemList, userList), nil
}

func getRelatedData(stx *repo.StoreTx, bidList []*m.BidModel) (
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
			return nil, nil, nil, errors.Next(err)
		}
		auctionList = append(auctionList, auction)
		auctionMap[auctionId] = auction
	}

	for _, auction := range auctionList {
		itemId := auction.ItemId
		item, err := stx.ItemRepo().GetById(itemId)
		if err != nil {
			return nil, nil, nil, errors.Next(err)
		}
		itemList = append(itemList, item)

		userId := auction.MaxBidderId
		user, err := stx.UserRepo().GetByProperty("id", userId)
		if err != nil {
			return nil, nil, nil, errors.Next(err)
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
) []*dtos.BidsTable {
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

	return bidsTable
}
