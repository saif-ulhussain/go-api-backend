package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/JSainsburyPLC/go-logrus-wrapper"
	. "go-api-backend/internal/models"
	"go-api-backend/internal/repository"
	"net/http"
)

type HabitHandler struct {
	habitRepository repository.HabitRepositoryInterface
}

func NewHabitHandler(habitRepository repository.HabitRepositoryInterface) *HabitHandler {
	return &HabitHandler{
		habitRepository: habitRepository,
	}
}

func (h *HabitHandler) CreateHabitHandler(w http.ResponseWriter, r *http.Request) {
	var habit Habit

	err := json.NewDecoder(r.Body).Decode(&habit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		log.Error(fmt.Sprintf("Invalid request body: %s", err))
		return
	}

	err = h.validateHabit(habit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		log.Error(fmt.Sprintf("Invalid request body: %s", err))
		return
	}

	err = h.habitRepository.InsertHabit(habit)
	if err != nil {
		http.Error(w, "Failed to insert new habit", http.StatusBadRequest)
		log.Error(fmt.Sprintf("Failed to insert new habit: %s", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Entry successfully created."))

	log.CtxInfof(r.Context(), "Entry successfully created.")
}

func (h *HabitHandler) validateHabit(habit Habit) error {
	if habit.Name == "" || habit.StartDate == "" {
		return errors.New("required properties are missing")
	}
	return nil
}
