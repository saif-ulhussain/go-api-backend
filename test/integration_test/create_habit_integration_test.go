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
	"net/http"
	"net/http/httptest"
)

func setupTestDatabase() (*sql.DB, error) {
	// Connect to the test database using the updated connection details
	db, err := sql.Open("postgres", "host=localhost port=5433 user=postgres password=mysecretpassword dbname=go-api-backend-db sslmode=disable")
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
		testDB *sql.DB
	)

	BeforeEach(func() {
		// Set up the test database
		var err error
		testDB, err = setupTestDatabase()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		// Clear the test database
		clearTestDatabase(testDB)
	})

	Context("When valid habit data is provided", func() {
		It("Should return a status response of 201", func() {
			habit := handlers.Habit{
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

			handlers.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusCreated))
		})
	})

	Context("When invalid habit data is provided", func() {
		It("Should return a status response of 400", func() {
			// Prepare the request body with missing required fields
			habit := handlers.Habit{
				Name: "Exercise 2",
			}
			habitJSON, err := json.Marshal(habit)
			Expect(err).NotTo(HaveOccurred())

			req := httptest.NewRequest("POST", "/create-habit", bytes.NewBuffer(habitJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			handlers.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("Failed to insert new habit"))
		})
	})
})
