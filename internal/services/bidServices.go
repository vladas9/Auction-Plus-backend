package services

import (
	"fmt"

	"github.com/google/uuid"
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
		return fmt.Errorf("Palacing bid faild: %s", err)
	}

	return nil
}

func (s *Service) ShowBidTable(userId uuid.UUID, limit, offset int) (responce []*m.BidsTable, err error) {
	var bidList []*m.BidModel
	var auctionList []*m.AuctionModel
	var itemList []*m.ItemModel

	err = s.store.WithTx(func(stx *r.StoreTx) error {
		bidList, err = stx.BidRepo().GetAllByUserId(userId, limit, offset)
		if err != nil {
			return fmt.Errorf("Failed geting bids: %s", err)
		}

		for _, bid := range bidList {
			auctionId := bid.AuctionId
			auction, err := stx.AuctionRepo().GetById(auctionId)

			if err != nil {
				return fmt.Errorf("Failed geting auction: %s", err)
			}

			auctionList = append(auctionList, auction)
		}

		for _, auction := range auctionList {
			itemId := auction.ItemId
			item, err := stx.ItemRepo().GetById(itemId)

			if err != nil {
				return fmt.Errorf("Failed geting item: %s", err)
			}

			itemList = append(itemList, item)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var bidsTable []*m.BidsTable

	auctionMap := make(map[uuid.UUID]*m.AuctionModel)
	for _, auction := range auctionList {
		auctionMap[auction.ID] = auction
	}

	itemMap := make(map[uuid.UUID]*m.ItemModel)
	for _, item := range itemList {
		itemMap[item.ID] = item
	}

	highestBids := make(map[uuid.UUID]*m.BidModel)

	for _, bid := range bidList {
		if existingBid, exists := highestBids[bid.AuctionId]; !exists || bid.Amount.Compare(existingBid.Amount) == 1 {
			highestBids[bid.AuctionId] = bid
		}
	}

	for _, bid := range highestBids {
		relatedAuction, auctionExists := auctionMap[bid.AuctionId]
		if !auctionExists {
			continue
		}

		relatedItem, itemExists := itemMap[relatedAuction.ItemId]
		if !itemExists {
			continue
		}

		var image uuid.UUID
		if len(relatedItem.Images) > 0 {
			image = relatedItem.Images[0]
		} else {
			image = uuid.Nil
		}

		bidTableEntry := m.BidsTableMapper(
			image,
			relatedAuction.CurrentBid,
			bid.Amount,
			Host,
			Port,
			relatedItem.Name,
			string(relatedItem.Category),
			relatedAuction.IsActive,
			relatedAuction.EndTime,
		)

		bidsTable = append(bidsTable, bidTableEntry)
	}

	return bidsTable, nil
}
