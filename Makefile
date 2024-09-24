BINARY_NAME=./bin/main
all: run
 
build:
	@go build -o ${BINARY_NAME} ./cmd/main.go
 
run: build
	${BINARY_NAME}

setup:
	@psql -U postgres -c "CREATE DATABASE auctiondb"
	@psql -U postgres -d auctiondb -f ./pkg/postgres/setup.sql
	# populating database
	@go run ./cmd/demodb

drop:
	psql -U postgres -d auctiondb -c "DROP DATABASE IF EXISTS auctiondb"

deps:
	@go mod tidy
