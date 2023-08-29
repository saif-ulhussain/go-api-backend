package integration_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/repository"
	"net/http"
	"net/http/httptest"
)

var testDB *sql.DB

var _ = Describe("LoginUserHandler", func() {
	var (
		userRepository *repository.UserRepository
		loginHandler   *handlers.LoginHandler
	)

	BeforeEach(func() {
		var err error
		testDB, err = setupTestDatabase()
		userRepository = repository.NewUserRepository(testDB)
		loginHandler = handlers.NewLoginHandler(userRepository)
		Expect(err).NotTo(HaveOccurred())
		seedTestData(testDB)
	})

	AfterEach(func() {
		clearTestDatabase(testDB)
	})

	Context("When valid user credential data is provided", func() {
		It("Should return a status code of 200", func() {
			loginData := struct {
				Email    string
				Password string
			}{
				Email:    "email@email.com",
				Password: "SecurePassword!",
			}

			loginDataJSON, err := json.Marshal(loginData)
			Expect(err).NotTo(HaveOccurred())

			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginDataJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			loginHandler.LoginUserHandler(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

		})
	})

	Context("When invalid user credential data is provided", func() {
		It("Should return a status code of 401", func() {
			invalidLoginData := struct {
				Email    string
				Password string
			}{
				Email:    "email@email.com",
				Password: "IncorrectPassword",
			}

			loginDataJSON, err := json.Marshal(invalidLoginData)
			Expect(err).NotTo(HaveOccurred())

			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(loginDataJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			loginHandler.LoginUserHandler(res, req)
			Expect(res.Code).To(Equal(http.StatusUnauthorized))

		})
	})
})
