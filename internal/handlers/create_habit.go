package handlers

import (
	"encoding/json"
	"fmt"
	"go-api-backend/internal/configuration"
	"go-api-backend/internal/models"
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
	var habit models.Habit
	err := json.NewDecoder(r.Body).Decode(&habit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.habitRepository.InsertHabit(habit)
	if err != nil {
		http.Error(w, "Failed to insert new habit", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Entry successfully created."))
	configuration.LogInfo("Entry successfully created.")
}
