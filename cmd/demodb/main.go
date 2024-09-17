package main

import (
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
	err = r.NewStore(db).WithTx(func(stx *r.StoreTx) error {
		//userId, err := stx.UserRepo().Insert(u.GenerateDummyUser())
		//if err != nil {
		//	return err
		//}
		userId, _ := uuid.Parse("da0db08a-0ab1-483c-a228-f592a8d43b8b")

		auctions := u.GenerateDummyAuctions(userId)
		for _, auct := range auctions {
			err := stx.AuctionRepo().Insert(auct)
			if err != nil {
				return err
			}
		}
		return nil

	})
	if err != nil {
		println(err.Error())
	}
}
