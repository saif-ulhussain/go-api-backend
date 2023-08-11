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

var _ = Describe("CreateUserHandler", func() {

	var userHandler *handlers.UserHandler
	var mockRepo *mocks.MockUserRepositoryInterface

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		defer ctrl.Finish()
		mockRepo = mocks.NewMockUserRepositoryInterface(ctrl)
		userHandler = handlers.NewUserHandler(mockRepo)
	})

	Context("When valid new user details is provided", func() {
		It("Should return a http status code of 201", func() {
			userJSON := []byte(`{
				"first_name": "S",
				"last_name": "Hussain",
				"email": "email@email.com"
			}`)

			mockRepo.EXPECT().InsertUser(gomock.Any()).Return(nil)

			req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userJSON))
			Expect(err).NotTo(HaveOccurred())

			res := httptest.NewRecorder()

			userHandler.CreateUserHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusCreated))
			Expect(res.Body.String()).To(Equal("User successfully created."))
		})
	})
})
