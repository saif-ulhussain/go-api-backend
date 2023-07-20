package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/models"
	"go-api-backend/mocks"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("CreateHabitHandler", func() {

	Context("When valid habit data is provided", func() {
		It("Should return a status response of 201", func() {

			ctrl := gomock.NewController(GinkgoT())
			defer ctrl.Finish()

			mockRepo := mocks.NewMockHabitRepositoryInterface(ctrl)
			habitHandler := handlers.NewHabitHandler(mockRepo)

			habit := models.Habit{
				Name:        "Exercise 3",
				StartDate:   "2012-01-02",
				EndDate:     nil,
				StreakCount: nil,
				Completed:   nil,
				Comments:    nil,
				Category:    nil,
			}

			mockRepo.EXPECT().InsertHabit(gomock.AssignableToTypeOf(habit)).Return(nil)

			habitJSON, err := json.Marshal(habit)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			req, err := http.NewRequest(http.MethodPost, "/habits", bytes.NewBuffer(habitJSON))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			res := httptest.NewRecorder()

			habitHandler.CreateHabitHandler(res, req)

			gomega.Expect(res.Code).To(gomega.Equal(http.StatusCreated))
			gomega.Expect(res.Body.String()).To(gomega.Equal("Entry successfully created."))
		})
	})
})
