package main

import (
	//"github.com/google/uuid"
	r "github.com/vladas9/backend-practice/internal/repository"
	u "github.com/vladas9/backend-practice/internal/utils"

	db "github.com/vladas9/backend-practice/pkg/postgres"
)

func main() {
	u.SetupLogger("./log-files/demodb.log")
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	user := u.GenerateDummyUser()
	err = r.NewStore(db).WithTx(func(stx *r.StoreTx) error {
		userId, repoErr := stx.UserRepo().Insert(user)
		u.Logger.Info("demoDB: userId: ", userId)
		if repoErr != nil {
			u.Logger.Error("demoDB stx.UserRepo():", repoErr)
			return repoErr
		}
		//userId, _ := uuid.Parse("da0db08a-0ab1-483c-a228-f592a8d43b8b")

		for i := 0; i <= 20; i++ {
			item := u.CreateDummyItem()
			itemId, err := stx.ItemRepo().Insert(item)
			if err != nil {
				return err
			}
			auct := u.CreateDummyAuction(itemId, userId)
			u.Logger.Info("\n", auct, "\n")
			err = stx.AuctionRepo().Insert(auct)
			if err != nil {
				return err
			}
		}
		return nil

	})
	if err != nil {
		u.Logger.Error("demoDB:", err)
	}
}
