package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var (
	host     = os.Getenv("HOST")
	port     = 5432
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	dbname   = os.Getenv("DBNAME")
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
