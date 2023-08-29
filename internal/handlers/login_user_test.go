package handlers_test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/models"
	"go-api-backend/mocks"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("LoginUserHandler", func() {

	var loginHandler *handlers.LoginHandler
	var mockRepo *mocks.MockUserRepositoryInterface

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		defer ctrl.Finish()
		mockRepo = mocks.NewMockUserRepositoryInterface(ctrl)
		loginHandler = handlers.NewLoginHandler(mockRepo)
	})

	Context("When valid user credentials are provided", func() {
		It("Should return a http status code of 200", func() {
			userJSON := []byte(`{
				"email": "email@email.com",
				"password": "SecurePassword!"
			}`)

			validUser := models.User{
				ID:        0,
				FirstName: "S",
				LastName:  "H",
				Email:     "email@email.com",
				Password:  "$2a$10$C8AyR0Yae1jNrzOUHuZRI.cs9wK/ZfpDJ3yQLXhUoX5iuyHwO/FJK",
			}

			mockRepo.EXPECT().GetUserByEmail(gomock.Any()).Return(validUser, nil)

			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJSON))
			Expect(err).NotTo(HaveOccurred())

			res := httptest.NewRecorder()

			loginHandler.LoginUserHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusOK))
			Expect(res.Body.String()).To(Equal("User successfully authenticated."))
		})
	})
})
