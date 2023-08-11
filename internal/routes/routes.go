package routes

import (
	"database/sql"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"github.com/JSainsburyPLC/go-logrus-wrapper/middleware"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	newrelicconfig "go-api-backend/internal/configuration/newrelic"
	"go-api-backend/internal/handlers"
	"go-api-backend/internal/repository"
	"net/http"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	routes := mux.NewRouter()
	routes.Use(middleware.AddLoggableHeadersToContext)

	app, err := newrelicconfig.Initialize()
	if err != nil {
		log.Error("Error initialising new relic. ", err)
	}

	habitRepository := repository.NewHabitRepository(db)
	userRepository := repository.NewUserRepository(db)
	habitHandler := handlers.NewHabitHandler(habitRepository)
	userHandler := handlers.NewUserHandler(userRepository)

	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/health", handlers.HealthCheckHandler)).Methods(http.MethodGet)
	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/habit", habitHandler.CreateHabitHandler)).Methods(http.MethodPost)
	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/user", userHandler.CreateUserHandler)).Methods(http.MethodPost)
	return routes
}
