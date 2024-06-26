package task

import (
	"database/sql"
	"fmt"
)

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       string  `json:"color"`
	StatusID    *string `json:"status_id"`
	UserID      *string `json:"user_id"`
}

type TaskForCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	// StatusID    string `json:"status_id"`
	// UserID      string `json:"user_id"`
}

type TaskForUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	// StatusID    string `json:"status_id"`
	// UserID      string `json:"user_id"`
}

var allColumns = []any{
	"id",
	"title",
	"description",
	"color",
	"status_id",
	"user_id",
}

func mapRowToTask(rows *sql.Rows) (Task, error) {
	var t Task
	err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Color, &t.StatusID, &t.UserID)

	if err != nil {
		return Task{}, fmt.Errorf("Error error scanning Task row: %w", err)

	}
	return t, nil
}
