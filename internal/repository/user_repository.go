//go:generate mockgen --destination=./mocks/mock_user_repository.go go-api-backend/internal/repository UserRepositoryInterface

package repository

import (
	"database/sql"
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	"go-api-backend/internal/models"
)

type UserRepositoryInterface interface {
	InsertUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) InsertUser(user models.User) error {
	query := "INSERT INTO \"user\" (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)"
	stmt, err := u.db.Prepare(query)

	if err != nil {
		log.Error(fmt.Sprintf("Insert User Error: %s", err))
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		log.Error(fmt.Sprintf("Insert User Error: %s", err))
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM \"user\"  WHERE email = $1"

	stmt, err := u.db.Prepare(query)
	if err != nil {
		log.Error(fmt.Sprintf("Get User By Email Error: %s", err))
		return user, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(fmt.Sprintf("GetUserByEmail Error: %s", err))
			return user, err
		}
		log.Error(fmt.Sprintf("GetUserByEmail Error: %s", err))
		return user, err
	}

	return user, nil
}
