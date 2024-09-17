package services

import (
	"fmt"

	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
)

type AuctionParams struct {
	Offset, Len int
}

func (a AuctionParams) Validate() Problems {
	if a.Offset > a.Len {
		return Problems{
			"limit":  "limmit smaller than offset",
			"offset": "offset bigger than limit",
		}
	}
	if a.Len <= 0 {
		return Problems{"limit": "must be more that 0"}
	}
	if a.Offset < 0 {
		return Problems{"offset": "cannot be negative"}
	}
	return nil
}

func (s *Service) GetAuctionData(params AuctionParams) ([]*m.AuctionModel, error) {
	var err error
	var list []*m.AuctionModel
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		list, err = stx.AuctionRepo().GetAll(params.Offset, params.Len)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("GetAuctions controller: %w", err)
	}
	return list, err
}

func (s *Service) CreateAuctions(auctions []*m.AuctionModel) error {
	err := s.store.WithTx(func(stx *r.StoreTx) error {
		for _, auct := range auctions {
			err := stx.AuctionRepo().Insert(auct)
			if err != nil {
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
