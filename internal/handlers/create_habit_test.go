package handlers_test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"go-api-backend/mocks"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("CreateHabitHandler", func() {

	var habitHandler *handlers.HabitHandler
	var mockRepo *mocks.MockHabitRepositoryInterface

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		defer ctrl.Finish()
		mockRepo = mocks.NewMockHabitRepositoryInterface(ctrl)
		habitHandler = handlers.NewHabitHandler(mockRepo)
	})

	Context("When valid habit data is provided", func() {
		It("Should return a status response of 201", func() {
			validJSON := []byte(`{
				"name": "Complete Leet Code Exercise...",
				"start_date": "2012-01-02",
				"end_date": null,
				"streak_count": 20,
				"completed": true,
				"comments": "great!",
				"category": "Golang"
			}`)

			mockRepo.EXPECT().InsertHabit(gomock.Any()).Return(nil)

			req, err := http.NewRequest(http.MethodPost, "/habits", bytes.NewBuffer(validJSON))
			Expect(err).NotTo(HaveOccurred())

			res := httptest.NewRecorder()

			habitHandler.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusCreated))
			Expect(res.Body.String()).To(Equal("Entry successfully created."))
		})
	})

	Context("When invalid habit data is provided / mandatory property is missing", func() {
		It("Should return a status response of 400", func() {
			invalidJSON := []byte(`{
            "name": "Complete Leet Code Exercise...",
            "start_date": null,
            "end_date": null,
            "streak_count": "IncorrectDataType"",
            "completed": true,
            "comments": "great!",
            "category": "Golang" }`)

			req, err := http.NewRequest(http.MethodPost, "/habits", bytes.NewBuffer(invalidJSON))
			Expect(err).NotTo(HaveOccurred())

			res := httptest.NewRecorder()

			mockRepo.EXPECT().InsertHabit(gomock.Any()).Return(nil)

			habitHandler.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("Invalid request body"))
		})
	})

	Context("When a mandatory property is missing", func() {
		It("Should return a status response of 400", func() {
			invalidJSON := []byte(`{
            "start_date": null,
            "end_date": null,
            "streak_count": 20,
            "completed": true,
            "comments": "great!",
            "category": "Golang" }`)

			req, err := http.NewRequest(http.MethodPost, "/habits", bytes.NewBuffer(invalidJSON))
			Expect(err).NotTo(HaveOccurred())

			res := httptest.NewRecorder()

			mockRepo.EXPECT().InsertHabit(gomock.Any()).Return(nil)

			habitHandler.CreateHabitHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("required properties are missing"))
		})
	})
})
