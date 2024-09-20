package services

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
)

func (s *Service) GetAuctions(params AuctionParams) (respList []*AuctionResp, err error) {
	var auctList []*m.AuctionModel
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		if auctList, err = getAuctionsWith(stx, params); err != nil {
			return err
		}
		for _, auct := range auctList {
			if auctResp, err := s.newAuctionResp(stx, auct).withItem().unpack(); err != nil {
				return err
			} else if auctResp.ItemHas(params.Condition, params.Category) {
				respList = append(respList, auctResp)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf(" Auction Service: GetAuctions: %w", err)
	}

	return respList, nil
}

func (s *Service) GetAuctionById(id uuid.UUID) (resp *AuctionResp, err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		auction, err := stx.AuctionRepo().GetById(id)
		if err != nil {
			return err
		}
		resp, err = s.newAuctionResp(stx, auction).withItem().withBids().unpack()
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("Auction Service: GetAuctionById: %w", err)
	}
	return resp, nil
}

func getAuctionsWith(stx *r.StoreTx, params AuctionParams) (auctions []*m.AuctionModel, err error) {
	auctionRepo := stx.AuctionRepo()
	if params.MaxPrice.IsZero() {
		auctions, err = auctionRepo.GetAll(params.Offset, params.Len)
	} else {
		auctions, err = auctionRepo.GetAllFiltered(
			params.Offset, params.Len,
			params.MinPrice, params.MaxPrice)
		u.Logger.Info("getAuctionsWith in:", auctions)
	}
	if err != nil {
		return nil, fmt.Errorf("getAuctionsWith: %w", err)
	}
	return auctions, nil
}

func (s *Service) CreateAuction(auct *m.AuctionModel, item *m.ItemModel) error {
	err := s.store.WithTx(func(stx *r.StoreTx) error {
		if itemId, err := stx.ItemRepo().Insert(item); err != nil {
			return err
		} else {
			auct.ItemId = itemId
		}
		if err := stx.AuctionRepo().Insert(auct); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("GetAuctions controller: %w", err)
	}
	return nil
}

//// Definitions

type AuctionParams struct {
	Category           string
	Condition          string
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
		return Problems{"max_price": "max price smaller than min price"}
	}
	if ok := m.IsCategory(a.Category); !ok {
		return Problems{"category:": fmt.Sprint(a.Category, "does not exist")}
	}
	if ok := m.IsCondition(a.Condition); !ok {
		return Problems{"condition:": fmt.Sprint(a.Condition, "does not exist")}
	}
	return nil
}

type AuctionResp struct {
	Auction *m.AuctionModel
	Item    *m.ItemModel
	BidList []*m.BidModel
}

func (rsp *AuctionResp) ItemHas(condition, category string) bool {
	hasCateg := (category == "" ||
		rsp.Item.Category == m.Category(category))
	hasCond := (condition == "" ||
		rsp.Item.Condition == m.Condition(condition))

	u.Logger.Info("condition:", hasCond, condition)
	u.Logger.Info("category:", hasCateg, category)

	if hasCond && hasCateg {
		return true
	}

	return false
}

type auctRespChain struct {
	err  error
	stx  *r.StoreTx
	Resp *AuctionResp
}

func (s *Service) newAuctionResp(stx *r.StoreTx, auct *m.AuctionModel) (next *auctRespChain) {
	next = &auctRespChain{
		err: nil,
		stx: stx,
		Resp: &AuctionResp{
			Auction: auct, Item: nil, BidList: nil,
		}}
	return
}

func (prev *auctRespChain) unpack() (*AuctionResp, error) {
	if prev.err != nil {
		return nil, fmt.Errorf("unpack: %w", prev.err)
	}
	u.Logger.Info("unpack: ", prev.Resp)
	return prev.Resp, nil
}

func (prev *auctRespChain) withBids() (next *auctRespChain) {
	next = &auctRespChain{}
	if prev.err != nil {
		next.err = fmt.Errorf("withBids prev: %w", prev.err)
		return
	}
	next = prev
	next.Resp.BidList, next.err = prev.stx.BidRepo().GetAllFor(prev.Resp.Auction)
	if next.err != nil {
		next.err = fmt.Errorf("withBids: %w", next.err)
	}
	return
}

func (prev *auctRespChain) withItem() (next *auctRespChain) {
	next = &auctRespChain{}
	if prev.err != nil {
		next.err = fmt.Errorf("withItem prev: %w", prev.err)
		return
	}
	next = prev
	next.Resp.Item, next.err = prev.stx.ItemRepo().GetById(prev.Resp.Auction.ItemId)
	if next.err != nil {
		next.err = fmt.Errorf("withItem: %w", next.err)
	}
	return
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
			auction.Status,
			auction.EndTime,
		)

		auctionTable = append(auctionTable, auctionTableEntry)
	}

	return auctionTable, nil
}
