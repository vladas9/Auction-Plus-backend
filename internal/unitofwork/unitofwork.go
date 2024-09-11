package unitofwork

import (
	"database/sql"
	//_ "github.com/lib/pq"
	//m "github.com/vladas9/backend-practice/internal/models"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"
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

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{
		db: db,
	}
}

// TODO handle edgecases: uncommitted tx
func (uow *UnitOfWork) BeginTransaction() error {
	if uow.tx != nil {
		if err := uow.tx.Rollback(); err != nil && err != sql.ErrTxDone {
			return err
		}
	}

	tx, err := uow.db.Begin()
	if err != nil {
		return err
	}

	uow.tx = tx
	uow.AuctionRepo = r.NewAuctionRepo(tx)
	uow.BidRepo = r.NewBidRepo(tx)
	uow.ItemRepo = r.NewItemRepo(tx)
	//NotificationRepo: r.NewNotificationRepo(tx),
	//ShippingRepo:     r.NewShippingRepo(tx),
	//TransactionRepo:  r.NewTransactionRepo(tx),
	uow.UserRepo = r.NewUserRepo(tx)
	return nil
}

func (uow *UnitOfWork) Commit() error {
	if uow.tx == nil {
		return nil
	}

	if err := uow.tx.Commit(); err != nil {
		u.Logger.Error("Failed to commit: ", err.Error())
		//uow.Rollback() ????
		return err
	}

	uow.tx = nil
	return nil
}

func (uow *UnitOfWork) Rollback() error {
	return uow.tx.Rollback()
}
