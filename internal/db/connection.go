package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-api-backend/internal/configuration"
	"os"
)

//var (
//	host     string
//	port     string
//	user     string
//	password string
//	dbname   string
//	//	host     = "localhost"
//	//	port     = 5432
//	//	user     = "postgres"
//	//	password = "mysecretpassword"
//	//	dbname   = "go-api-backend-db"
//)

const (
	connString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func GetConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	return fmt.Sprintf(connString, host, port, user, password, dbname)
}

func ConnectToDB() (*sql.DB, error) {
	configuration.LogInfo(fmt.Sprintf("Database Attempting to Connect."))

	db, err := sql.Open("postgres", GetConnectionString())
	if err != nil {
		configuration.LogError(fmt.Sprintf("Database connection error: %s", err))

		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	configuration.LogInfo("Successfully connected to database.")
	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
