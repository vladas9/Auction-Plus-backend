package main

import (
	//"github.com/google/uuid"
	"github.com/google/uuid"
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
	//user := GenerateDummyUser()
	err = r.NewStore(db).WithTx(func(stx *r.StoreTx) error {
		//userId, repoErr := stx.UserRepo().Insert(user)
		//u.Logger.Info("demoDB: userId: ", userId)
		//if repoErr != nil {
		//	u.Logger.Error("demoDB stx.UserRepo():", repoErr)
		//	return repoErr
		//}
		userId, _ := uuid.Parse("4d3b4f40-a35d-44ff-867e-487fc29911ca")

		for i := 0; i <= 20; i++ {
			item := CreateDummyItem()
			itemId, err := stx.ItemRepo().Insert(item)
			if err != nil {
				return err
			}
			auct := CreateDummyAuction(itemId, userId)
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
