package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/repository"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	routes := mux.NewRouter()

	app, err := newreliconfig.Initialize()

	if err != nil {
		LogInfo(err)
	}

	habitRepository := repository.NewHabitRepository(db)
	habitHandler := handlers.NewHabitHandler(habitRepository)

	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/health", handlers.HealthCheckHandler)).Methods("GET")
	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/create-habit", habitHandler.CreateHabitHandler)).Methods("POST")
	return routes
}
