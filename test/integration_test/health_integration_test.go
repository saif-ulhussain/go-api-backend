package integration

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-api-backend/internal/handlers"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Health Check", func() {
	Context("Checks health status", func() {
		It("Should return a status response of 200", func() {
			req := httptest.NewRequest("GET", "/health", nil)
			res := httptest.NewRecorder()

			handlers.HealthCheckHandler(res, req)

			Expect(res.Code).To(Equal(http.StatusOK))
			Expect(res.Body.String()).To(Equal("OK"))
		})
	})
})
