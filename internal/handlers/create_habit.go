package handlers

import (
	"encoding/json"
	"fmt"
	"go-api-backend/internal/db"
	"log"
	"net/http"
)

type Habit struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	StreakCount *int    `json:"streak_count"`
	Completed   *bool   `json:"completed"`
	Comments    *string `json:"comments"`
	Category    *string `json:"category"`
}

func CreateHabitHandler(w http.ResponseWriter, r *http.Request) {
	var habit Habit
	err := json.NewDecoder(r.Body).Decode(&habit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = InsertHabit(habit)
	if err != nil {
		http.Error(w, "Failed to insert new habit", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Entry successfully created."))
}

func InsertHabit(habit Habit) error {
	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "INSERT INTO habit (name, start_date, end_date, streak_count, completed, comments, category) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = db.Exec(query, habit.Name, habit.StartDate, habit.EndDate, habit.StreakCount, habit.Completed, habit.Comments, habit.Category)

	if err != nil {
		return err
	}

	return nil
}
