package main

import (
	"fmt"
	"go-api-backend/internal/routes"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println(err)
	}
}
