package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	. "go-api-backend/internal/models"
	"go-api-backend/internal/repository"
	"net/http"
)

type UserHandler struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserHandler(userRepository repository.UserRepositoryInterface) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		log.Error(fmt.Sprintf("Invalid request body: %s", err))
		return
	}

	err = u.userRepository.InsertUser(user)
	if err != nil {
		http.Error(w, "Failed to insert new user", http.StatusBadRequest)
		log.Error(fmt.Sprintf("Failed to insert new user: %s", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User successfully created."))

	log.CtxInfof(r.Context(), "User successfully created.")
}