package utilities

import (
	"context"
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

func GenerateJWT(userID int) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["user"] = userID

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("JWT_SECRET_KEY")
		tokenHeader := r.Header.Get("Token")
		if tokenHeader == "" {
			http.Error(w, "Token missing", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			log.Error(fmt.Sprintf("Invalid token: %s", err))
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Error(fmt.Sprintf("Unable to extract claims: %s", err))
			http.Error(w, "Unable to extract claims", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "JWT", claims)
		next(w, r.WithContext(ctx))
	})
}
