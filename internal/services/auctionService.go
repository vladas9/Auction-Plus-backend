package services

import (
	"fmt"

	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	"github.com/vladas9/backend-practice/internal/utils"
)

type AuctionParams struct {
	Category           m.Category
	LotState           m.Condition
	Offset, Len        int
	MinPrice, MaxPrice m.Decimal
}

func (a AuctionParams) Validate() Problems {
	if a.Len <= 0 {
		return Problems{"limit": "must be more that 0"}
	}
	if a.Offset < 0 {
		return Problems{"offset": "cannot be negative"}
	}
	if !a.MaxPrice.IsZero() &&
		a.MaxPrice.Compare(a.MinPrice) == -1 {
		return Problems{"filters": "max price smaller than min price"}
	}
	return nil
}

func (s *Service) GetAuctions(params AuctionParams) (auctList []*m.AuctionModel, itemList []*m.ItemModel, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {

		if auctList, err = getAuctionsTx(stx, params); err != nil {
			return fmt.Errorf("getAuctionsTx error: %w", err)
		}
		for _, auct := range auctList {
			item, err := stx.ItemRepo().GetById(auct.ItemId)
			if err != nil {
				return err
			}
			itemList = append(itemList, item)
		}
		return err
	})

	if err != nil {
		return nil, nil, fmt.Errorf("GetAuctions controller: %w", err)
	}

	utils.Logger.Info("getAuctions:", auctList, itemList)
	return auctList, itemList, err
}

func getAuctionsTx(stx *r.StoreTx, params AuctionParams) (auctions []*m.AuctionModel, err error) {
	auctionRepo := stx.AuctionRepo()
	if params.MaxPrice.IsZero() {
		auctions, err = auctionRepo.GetAll(params.Offset, params.Len)
	} else {
		auctions, err = auctionRepo.GetAllFiltered(
			params.Offset, params.Len,
			params.MinPrice, params.MaxPrice)
		utils.Logger.Info("getAuctionsTx in:", auctions)
	}
	if err != nil {
		auctions = nil
	}
	return auctions, err
}

func (s *Service) CreateAuctions(auctions []*m.AuctionModel, items []*m.ItemModel) error {
	err := s.store.WithTx(func(stx *r.StoreTx) error {
		for i, auct := range auctions {
			if itemId, err := stx.ItemRepo().Insert(items[i]); err != nil {
				return err
			} else {
				auct.ItemId = itemId
			}
			if err := stx.AuctionRepo().Insert(auct); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("GetAuctions controller: %w", err)
	}
	return nil
}
