package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	u "github.com/vladas9/backend-practice/internal/utils"
)

var DB *sql.DB

func ConnectDB() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	connectStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	DB, err = sql.Open("postgres", connectStr)
	if err != nil {
		return err
	}

	err = DB.Ping()

	if err != nil {
		return err
	}

	u.Logger.Info("Successfully connected to the database")
	return nil
}
