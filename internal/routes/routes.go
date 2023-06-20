package routes

import (
	"github.com/gorilla/mux"
	"go-api-backend/internal/handlers"
)

func SetupRoutes() *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	return routes
}
