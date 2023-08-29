package models

type Habit struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	StreakCount *int    `json:"streak_count"`
	Completed   *bool   `json:"completed"`
	Comments    *string `json:"comments"`
	Category    *string `json:"category"`
	UserID      int
}
