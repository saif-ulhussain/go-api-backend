package integration_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/models"
	"go-api-backend/internal/repository"
	"net/http"
	"net/http/httptest"
)

func setupTestDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5433 user=postgres password=mysecretpassword dbname=go-api-backend-db-test sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the test db: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping the test db: %v", err)
	}

	return db, nil
}

func clearTestDatabase(db *sql.DB) {
	_, _ = db.Exec("TRUNCATE TABLE habit")
}

var _ = Describe("CreateHabitHandler", func() {
	var (
		testDB          *sql.DB
		habitRepository *repository.HabitRepository
		habitHandler    *handlers.HabitHandler
	)

	BeforeEach(func() {
		var err error
		testDB, err = setupTestDatabase()
		habitRepository = repository.NewHabitRepository(testDB)
		habitHandler = handlers.NewHabitHandler(habitRepository)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		clearTestDatabase(testDB)
	})

	Context("When valid habit data is provided", func() {
		It("Should return a status response of 201", func() {
			habit := models.Habit{
				Name:        "Exercise 3",
				StartDate:   "2012-01-02",
				EndDate:     nil,
				StreakCount: nil,
				Completed:   nil,
				Comments:    nil,
				Category:    nil,
			}
			habitJSON, err := json.Marshal(habit)
			Expect(err).NotTo(HaveOccurred())

			req := httptest.NewRequest("POST", "/create-habit", bytes.NewBuffer(habitJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			habitHandler.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusCreated))
		})
	})

	Context("When invalid habit data is provided", func() {
		It("Should return a status response of 400", func() {
			// Prepare the request body with missing required fields
			habit := models.Habit{
				Name: "Exercise 2",
			}
			habitJSON, err := json.Marshal(habit)
			Expect(err).NotTo(HaveOccurred())

			req := httptest.NewRequest("POST", "/create-habit", bytes.NewBuffer(habitJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			habitHandler.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("Invalid request body"))
		})
	})
})
