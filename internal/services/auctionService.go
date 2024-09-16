package services

import (
	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
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
