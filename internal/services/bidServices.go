package services

import (
	"fmt"

	m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
)

func (s *Service) ValidateBid() {

}

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
