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
	"go-api-backend/internal/utilities"
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
	userHandler := handlers.NewRegisterHandler(userRepository)
	loginHandler := handlers.NewLoginHandler(userRepository)

	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/health", handlers.HealthCheckHandler)).Methods(http.MethodGet)
	routes.Handle(newrelic.WrapHandle(app, "/habit", utilities.ValidateJWT(habitHandler.CreateHabitHandler))).Methods(http.MethodPost)
	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/register", userHandler.RegisterUserHandler)).Methods(http.MethodPost)
	routes.HandleFunc(newrelic.WrapHandleFunc(app, "/login", loginHandler.LoginUserHandler)).Methods(http.MethodPost)
	return routes
}
