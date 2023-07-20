//go:generate mockgen -destination=../mocks/mock_habit_repository.go -package=mocks go-api-backend/internal/repository HabitRepository

package repository

import (
	"database/sql"
	"go-api-backend/internal/models"
)

type HabitRepository struct {
	db *sql.DB
}

type HabitRepositoryInterface interface {
	InsertHabit(habit models.Habit) error
}

func NewHabitRepository(db *sql.DB) *HabitRepository {
	return &HabitRepository{db: db}
}

func (r *HabitRepository) InsertHabit(habit models.Habit) error {
	query := "INSERT INTO habit (name, start_date, end_date, streak_count, completed, comments, category) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(query, habit.Name, habit.StartDate, habit.EndDate, habit.StreakCount, habit.Completed, habit.Comments, habit.Category)

	if err != nil {
		return err
	}

	return nil
}
