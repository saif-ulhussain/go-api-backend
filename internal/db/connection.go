package db

import (
	"database/sql"
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	_ "github.com/lib/pq"
	"os"
)

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
	log.Info("Database Attempting to Connect.")

	db, err := sql.Open("postgres", GetConnectionString())
	if err != nil {
		log.Error(fmt.Sprintf("Database connection error: %s", err))
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	log.Info("Successfully connected to database.")
	return db, nil
}
