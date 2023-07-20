package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go-api-backend/internal/configuration"
	"go-api-backend/internal/db"
	"go-api-backend/internal/routes"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	log.SetOutput(logger.Writer())
	configuration.InitialiseLogger(logger)

	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := routes.SetupRoutes(db)
	err = http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println(err)
	}
}
