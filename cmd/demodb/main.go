package main

import (
	"github.com/vladas9/backend-practice/internal/controllers"
	s "github.com/vladas9/backend-practice/internal/services"
	u "github.com/vladas9/backend-practice/internal/utils"

	db "github.com/vladas9/backend-practice/pkg/postgres"
)

func main() {
	u.SetupLogger("./log-files/demodb.log")
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ctrls := controllers.ControllersWith(db)
	err = ctrls.WithTx(func(service *s.Service) error {
		auctions := u.GenerateDummyAuctions()
		return service.CreateAuctions(auctions)
	})
	if err != nil {
		println(err.Error())
	}
}
