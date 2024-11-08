BINARY_NAME=./bin/main
all: run
 
build:
	@go build -o ${BINARY_NAME} ./cmd/main.go
 
run: build
	@#rm -f ./log-files/logs.log
	@echo "Starting server, run \`make log\` to see logs"
	@${BINARY_NAME}

log:
	@echo "Server Logs:"
	@tail -f ./log-files/logs.log

setup:
	@psql -U postgres -c "CREATE DATABASE auctiondb"
	@psql -U postgres -d auctiondb -f ./pkg/postgres/setup.sql
	@# populating database
	@go run ./cmd/demodb

drop:
	psql -U postgres -c "DROP DATABASE IF EXISTS auctiondb"

deps:
	@go mod tidy
