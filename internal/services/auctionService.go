package services

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
)

type AuctionParams struct {
	Offset, Len int
}

func (s *Service) GetAuctions(params AuctionParams) (list []*m.AuctionModel, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		list, err = stx.AuctionRepo().GetAll(params.Offset, params.Offset*params.Len)
		return err
	})
	return list, err
}

func (s *Service) CreateAuctions(auctions []*m.AuctionModel) (err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		for _, auct := range auctions {
			err := stx.AuctionRepo().Insert(auct)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (s *Service) GetAuctionTable(userId uuid.UUID, limit, offset int) ([]*m.AuctionTable, error) {
	var auctionList []*m.AuctionModel
	var itemList []*m.ItemModel
	var userList []*m.UserModel

	err := s.store.WithTx(func(stx *r.StoreTx) error {
		var err error
		auctionList, err = stx.AuctionRepo().GetAllByUserId(userId, limit, offset)
		if err != nil {
			return fmt.Errorf("Failed getting acution list: %s", err)
		}

		itemList, userList, err = getRelatedAuctionTableData(stx, auctionList)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	auctionTable, err := buildAuctionTable(auctionList, itemList, userList)
	if err != nil {
		return nil, err
	}

	return auctionTable, nil

}

func getRelatedAuctionTableData(stx *r.StoreTx, auctionList []*m.AuctionModel) (
	[]*m.ItemModel, []*m.UserModel, error,
) {
	itemList := make([]*m.ItemModel, 0)
	userList := make([]*m.UserModel, 0)

	for _, auction := range auctionList {
		itemId := auction.ItemId
		item, err := stx.ItemRepo().GetById(itemId)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed getting item: %s", err)
		}
		itemList = append(itemList, item)

		userId := auction.MaxBidderId
		user, err := stx.UserRepo().GetById(userId)
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, nil, fmt.Errorf("Failed getting user: %s", err)
			}
			user = &m.UserModel{}
		}
		userList = append(userList, user)
		userList = append(userList, user)
	}

	return itemList, userList, nil
}

func buildAuctionTable(
	auctionList []*m.AuctionModel,
	itemList []*m.ItemModel,
	userList []*m.UserModel,
) ([]*m.AuctionTable, error) {
	userMap := u.CreateUserMap(userList)
	auctionMap := u.CreateAuctionMap(auctionList)
	itemMap := u.CreateItemMap(itemList)

	auctionTable := make([]*m.AuctionTable, 0)
	for _, auction := range auctionMap {
		relatedItem, itemExists := itemMap[auction.ItemId]
		if !itemExists {
			continue
		}

		relatedUser, userExists := userMap[auction.MaxBidderId]
		if !userExists {
			continue
		}

		image := u.GetFirstImageOrNil(relatedItem)

		auctionTableEntry := m.AuctionTableMapper(
			image,
			auction.CurrentBid,
			Host,
			Port,
			relatedItem.Name,
			string(relatedItem.Category),
			relatedUser.Username,
			auction.IsActive,
			auction.EndTime,
		)

		auctionTable = append(auctionTable, auctionTableEntry)
	}

	return auctionTable, nil
}
