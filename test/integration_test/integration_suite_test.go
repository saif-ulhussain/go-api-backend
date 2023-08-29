package integration_test

import (
	"database/sql"
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"github.com/joho/godotenv"
	"go-api-backend/internal/models"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	db sql.DB
)

const (
	connString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func getConnectionString() string {
	host := os.Getenv("DB_HOST_TEST")
	port := os.Getenv("DB_PORT_TEST")
	user := os.Getenv("DB_USER_TEST")
	password := os.Getenv("DB_PASSWORD_TEST")
	dbname := os.Getenv("DB_NAME_TEST")
	return fmt.Sprintf(connString, host, port, user, password, dbname)
}

func setupTestDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the test db: %v", err)
	}

	db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping the test db: %v", err)
	}

	return db, nil
}

func clearTestDatabase(db *sql.DB) {
	db.Exec("TRUNCATE TABLE habit, \"user\"")
}

func TestIntegration(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file:% s", err)
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func seedTestData(db *sql.DB) {
	userData := models.User{
		ID:        1,
		FirstName: "S",
		LastName:  "H",
		Email:     "email@email.com",
		Password:  "$2a$10$C8AyR0Yae1jNrzOUHuZRI.cs9wK/ZfpDJ3yQLXhUoX5iuyHwO/FJK",
	}

	query := "INSERT INTO \"user\" (id, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)"
	stmt, err := db.Prepare(query)

	if err != nil {
		log.Error(fmt.Sprintf("Insert User Error: %s", err))
	}

	defer stmt.Close()

	_, err = stmt.Exec(userData.ID, userData.FirstName, userData.LastName, userData.Email, userData.Password)

	if err != nil {
		log.Error(fmt.Sprintf("Insert User Error: %s", err))
	}
}
