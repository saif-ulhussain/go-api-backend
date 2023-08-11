package integration_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/models"
	"go-api-backend/internal/repository"
	"net/http"
	"net/http/httptest"
)

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

			req := httptest.NewRequest("POST", "/habit", bytes.NewBuffer(habitJSON))
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

			req := httptest.NewRequest("POST", "/habit", bytes.NewBuffer(habitJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			habitHandler.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("Invalid request body"))
		})
	})
})
