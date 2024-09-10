package unitofwork

import (
	"database/sql"
	//_ "github.com/lib/pq"
	//m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
)

type UnitOfWork struct {
	db *sql.DB
	tx *sql.Tx

	AuctionRepo      *r.AuctionRepo
	BidRepo          *r.BidRepo
	ItemRepo         *r.ItemRepo
	NotificationRepo *r.NotificationRepo
	ShippingRepo     *r.ShippingRepo
	TransactionRepo  *r.TransactionRepo
	UserRepo         *r.UserRepo
}

func NewUnitOfWork(db *sql.DB) (*UnitOfWork, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &UnitOfWork{
		db:               db,
		tx:               tx,
		AuctionRepo:      r.NewAuctionRepo(tx),
		BidRepo:          r.NewBidRepo(tx),
		ItemRepo:         r.NewItemRepo(tx),
		NotificationRepo: r.NewNotificationRepo(tx),
		ShippingRepo:     r.NewShippingRepo(tx),
		TransactionRepo:  r.NewTransactionRepo(tx),
		UserRepo:         r.NewUserRepo(tx),
	}, nil
}

func (uow *UnitOfWork) Commit() error {
	return uow.tx.Commit()
}

func (uow *UnitOfWork) Rollback() error {
	return uow.tx.Rollback()
}
