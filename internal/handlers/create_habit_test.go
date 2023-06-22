package handlers_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("CreateHabitHandler", func() {
	Context("When valid habit data is provided", func() {
		It("Should return a status response of 201", func() {
			habitJSON := []byte(`{
				"name": "Complete Leet Code Exercise...",
				"start_date": "2012-01-02",
				"end_date": null,
				"streak_count": 20,
				"completed": true,
				"comments": "great!",
				"category": "Golang"
			}`)

			req := httptest.NewRequest("POST", "/create-habit", bytes.NewBuffer(habitJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			handlers.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusCreated))
		})
	})

	Context("When invalid habit data is provided, such as streak count provided in string format", func() {
		It("Should return a status response of 400", func() {
			habitJSON := []byte(`{
				"name": "Complete Leet Code Exercise...",
				"start_date": "2012-01-02",
				"end_date": null,
				"streak_count": "20"",
				"completed": true,
				"comments": "great!",
				"category": "Golang"
			}`)

			req := httptest.NewRequest("POST", "/create-habit", bytes.NewBuffer(habitJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			handlers.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("Invalid request body"))
		})
	})

	Context("When incomplete habit data is provided / mandatory property is missing", func() {
		It("Should return a status response of 400", func() {
			// Invalid habit data: missing required fields
			habitJSON := []byte(`{
				"name": "Complete Leet Code Exercise...",
				"end_date": null,
				"streak_count": 20,
				"completed": true,
				"comments": "great!",
				"category": "Golang"
			}`)

			req := httptest.NewRequest("POST", "/create-habit", bytes.NewBuffer(habitJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			handlers.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("Failed to insert new habit"))
		})
	})
})
