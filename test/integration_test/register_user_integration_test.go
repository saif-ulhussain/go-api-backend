package integration_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/models"
	"go-api-backend/internal/repository"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("RegisterUserHandler", func() {
	var (
		testDB          *sql.DB
		userRepository  *repository.UserRepository
		registerHandler *handlers.RegisterHandler
	)

	BeforeEach(func() {
		var err error
		testDB, err = setupTestDatabase()
		userRepository = repository.NewUserRepository(testDB)
		registerHandler = handlers.NewRegisterHandler(userRepository)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		clearTestDatabase(testDB)
	})

	Context("When valid user data is provided", func() {
		It("Should return a status code of 201", func() {
			user := models.User{
				FirstName: "S",
				LastName:  "H",
				Email:     "email@email.com",
				Password:  "P@ssword!",
			}

			userJSON, err := json.Marshal(user)
			Expect(err).NotTo(HaveOccurred())

			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			registerHandler.RegisterUserHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusCreated))

		})
	})
	Context("When invalid user data is provided", func() {
		It("Should return a status response of 400", func() {
			user := models.User{
				FirstName: "S",
			}

			userJson, err := json.Marshal(user)
			Expect(err).NotTo(HaveOccurred())

			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJson))
			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			registerHandler.RegisterUserHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusBadRequest))
			Expect(res.Body.String()).To(ContainSubstring("Invalid request body"))
		})
	})
})
