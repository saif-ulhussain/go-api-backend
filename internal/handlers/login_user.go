package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"go-api-backend/internal/repository"
	"go-api-backend/internal/utilities"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginHandler struct {
	userRepository repository.UserRepositoryInterface
}

func NewLoginHandler(userRepository repository.UserRepositoryInterface) *LoginHandler {
	return &LoginHandler{
		userRepository: userRepository,
	}
}

func (l *LoginHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		log.Error(fmt.Sprintf("Invalid request body: %s", err))
		return
	}

	user, err := l.userRepository.GetUserByEmail(loginData.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid email or password: %s", err.Error()), http.StatusUnauthorized)
		log.Error(fmt.Sprintf("Invalid email or password: %s", err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid email or password: %s", err.Error()), http.StatusUnauthorized)
		log.Error(fmt.Sprintf("Invalid email or password: %s", err))
		return
	}

	jwtToken, err := utilities.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to generate JWT: %s", err.Error()), http.StatusInternalServerError)
		log.Error(fmt.Sprintf("Unable to generate JWT: %s", err))
		return
	}

	cookie := http.Cookie{
		Name:     "Token",
		Value:    jwtToken,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully authenticated."))
	log.CtxInfof(r.Context(), "User successfully authenticated.")
}
