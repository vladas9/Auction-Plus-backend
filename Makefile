run:
	go run ./cmd/main.go

setupdb:
	psql -U postgres -d auctiondb -f ./pkg/setup.sql
