package services

import (
	"fmt"

	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
)

func (s *Service) NewBid(bid *m.BidModel) (err error) {
	err = s.store.WithTx(func(stx *r.StoreTx) error {
		err = stx.BidRepo().Insert(bid)
		return err
	})
	if err != nil {
		return fmt.Errorf("Palacing bid faild: %s", err)
	}

	return nil
}
