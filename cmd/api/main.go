package main

import (
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"github.com/joho/godotenv"
	"go-api-backend/internal/db"
	"go-api-backend/internal/routes"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	log.Enable("info")

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
