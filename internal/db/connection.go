package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "go-api-backend-db"
)

const (
	connString = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

func GetConnectionString() string {

	return fmt.Sprintf(connString, host, port, user, password, dbname)
}

func ConnectToDB() (*sql.DB, error) {
	godotenv.Load()
	db, err := sql.Open("postgres", GetConnectionString())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
