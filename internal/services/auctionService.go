package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	dto "github.com/vladas9/backend-practice/internal/dtos"
	"github.com/vladas9/backend-practice/internal/errors"
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
)

type AuctionService interface {
	NewAuction(dto *dto.AuctionFull) error
	GetAuctionCards(params AuctionCardParams) ([]dto.AuctionCard, error)
	GetFullAuctionById(id uuid.UUID) (*dto.AuctionFull, error)
	GetAuctionTable(params AuctionTableParams) ([]dto.AuctionTable, error)
}

// type auctionService struct{ *Service }
//
//	func NewAuctionService(s *Service) AuctionService {
//		return &auctionService{s}
//	}

type AuctionCardParams struct {
	Category           string
	Condition          string
	Offset, Len        int
	MinPrice, MaxPrice m.Decimal
}

func (a AuctionCardParams) Validate() Problems {
	problems := Problems{}

	if a.Len < 0 {
		problems["len"] = "cannot be negative"
	}
	if a.Offset < 0 {
		problems["offset"] = "cannot be negative"
	}
	if a.MinPrice.IsNegative() {
		problems["min_price"] = "cannot be nebative"
	}
	if a.MaxPrice.IsNegative() {
		problems["max_price"] = "cannot be nebative"
	} else if !a.MaxPrice.IsZero() && a.MaxPrice.Compare(a.MinPrice) == -1 {
		problems["max_price"] = "max price cannot be less than min price"
	}
	if ok := m.IsCategory(a.Category); !ok {
		problems["category"] = fmt.Sprintf("%s does not exist", a.Category)
	}
	if ok := m.IsCondition(a.Condition); !ok {
		problems["condition"] = fmt.Sprintf("%s does not exist", a.Condition)
	}

	if len(problems) > 0 {
		return problems
	}
	return nil
}

type AuctionTableParams struct {
	UserId        uuid.UUID
	Limit, Offset int
}

func (p AuctionTableParams) Validate() Problems {
	problems := Problems{}
	u.Logger.Info(p.Limit)
	if p.Limit <= 0 {
		problems["limit"] = "must be greater than 0"
	}
	if p.Offset < 0 {
		problems["offset"] = "cannot be negative"
	}

	if len(problems) > 0 {
		return problems
	}
	return nil
}

func (s *Service) NewAuction(dto *dto.AuctionFull, seller uuid.UUID) (uuid.UUID, error) {
	images := []uuid.UUID{}
	var auctId uuid.UUID
	var err error

	for _, img := range dto.ImgSrc {
		id := uuid.New()
		if err := u.DecodeAndSaveImage(img, ImageDir, id.String()); err != nil {
			return uuid.UUID{}, errors.Internal(err)
		}
		images = append(images, id)
	}

	item := &m.ItemModel{
		Name:        dto.Title,
		Description: dto.Description,
		Category:    dto.Category,
		Condition:   dto.Condition,
		Images:      images,
	}
	auct := &m.AuctionModel{
		StartPrice: dto.StartPrice,
		StartTime:  time.Now(),
		EndTime:    dto.EndDate,
		SellerId:   seller,
	}
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		if itemId, err := stx.ItemRepo().Insert(item); err != nil {
			return err
		} else {
			auct.ItemId = itemId
		}
		if auctId, err = stx.AuctionRepo().Insert(auct); err != nil {
			return err
		}
		return nil
	})
	return auctId, errors.Internal(err)
}

func (s *Service) GetAuctionTable(params AuctionTableParams) ([]dto.AuctionTable, error) {
	if problems := params.Validate(); problems != nil {
		return nil, problems.toErr()
	}
	var auctions []*m.AuctionDetails
	var err error
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		auctions, err = getUserAuctions(stx, params, withItem, withMaxBidder)
		return err
	})
	if err != nil {
		return nil, errors.Internal(err)
	}
	var table []dto.AuctionTable
	for _, a := range auctions {
		table = append(table, *dto.MapAuctionTable(a))
	}
	return table, nil
}

func getUserAuctions(stx *r.StoreTx, params AuctionTableParams, opts ...auctOpt) (auctions []*m.AuctionDetails, err error) {
	var auctModels []*m.AuctionModel
	auctModels, err = stx.AuctionRepo().GetAllByUserId(params.UserId, params.Limit, params.Offset)
	var auct *m.AuctionDetails
	for _, auctModel := range auctModels {
		auct, err = getAuctionDetails(auctModel, stx, opts...)
		if err != nil {
			return nil, errors.Next(err)
		}
		auctions = append(auctions, auct)
	}
	return auctions, nil
}

func (s *Service) GetAuctionCards(params AuctionCardParams) ([]dto.AuctionCard, error) {
	if problems := params.Validate(); problems != nil {
		return nil, problems.toErr()
	}
	var err error
	var auctDetails []*m.AuctionDetails
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		if auctDetails, err = getAuctions(stx, params, withItem); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, errors.Internal(err)
	}

	var cards []dto.AuctionCard
	for _, a := range auctDetails {
		cards = append(cards, *dto.MapAuctionCard(a))
	}
	return cards, nil
}

func getAuctions(stx *r.StoreTx, params AuctionCardParams, opts ...auctOpt) (auctions []*m.AuctionDetails, err error) {
	auctionRepo := stx.AuctionRepo()
	var auctModels []*m.AuctionModel
	auctModels, err = auctionRepo.GetAllFiltered(
		params.Offset, params.Len,
		params.MinPrice, params.MaxPrice)
	u.Logger.Info("getAuctionsWith in:", auctions)
	if err != nil {
		return nil, errors.Next(err)
	}
	for _, auctModel := range auctModels {
		auct, err := getAuctionDetails(auctModel, stx, opts...)
		if err != nil {
			return nil, errors.Next(err)
		} else if auct.ItemHas(params.Condition, params.Category) {
			auctions = append(auctions, auct)
		}
	}
	return auctions, nil
}

func (s *Service) GetFullAuctionById(id uuid.UUID) (*dto.AuctionFull, error) {
	var err error
	var auct *m.AuctionDetails
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		auctModel, err := stx.AuctionRepo().GetById(id)
		if err != nil {
			return errors.NotFound("auction not found", err)
		}
		auct, err = getAuctionDetails(auctModel, stx, withItem, withBids)
		return errors.Next(err)
	})
	if err != nil {
		return nil, err
	}
	return dto.MapAuctionRespToFull(auct), nil
}

type auctOpt func(stx *r.StoreTx, auct *m.AuctionDetails) (*m.AuctionDetails, error)

func getAuctionDetails(auct *m.AuctionModel, stx *r.StoreTx, opts ...auctOpt) (*m.AuctionDetails, error) {
	var err error
	details := m.NewAuctionDetails(auct)
	for _, opt := range opts {
		details, err = opt(stx, details)
		if err != nil {
			return nil, errors.Next(err)
		}
	}
	return details, nil
}

func withBids(stx *r.StoreTx, auct *m.AuctionDetails) (*m.AuctionDetails, error) {
	var err error
	auct.BidList, err = stx.BidRepo().GetAllFor(auct.Auction)
	if err != nil {
		return nil, errors.Next(err)
	}
	return auct, nil
}

func withItem(stx *r.StoreTx, auct *m.AuctionDetails) (*m.AuctionDetails, error) {
	var err error
	auct.Item, err = stx.ItemRepo().GetById(auct.Auction.ItemId)
	if err != nil {
		return nil, errors.Next(err)
	}
	return auct, nil
}

func withMaxBidder(stx *r.StoreTx, auct *m.AuctionDetails) (*m.AuctionDetails, error) {
	var err error
	auct.MaxBidder, err = stx.UserRepo().GetByProperty("id", auct.Auction.MaxBidderId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Next(err)
		}
		auct.MaxBidder = &m.UserModel{}
	}
	return auct, nil
}
