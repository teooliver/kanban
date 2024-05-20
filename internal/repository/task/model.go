package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/doug-martin/goqu"
)

type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	StatusID    *string `json:"status_id"`
	Color       *string `json:"color"`
	UserID      *string `json:"user_id"`
}

type TaskForCreate struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	StatusID    *string `json:"status_id"`
	Color       *string `json:"color"`
	UserID      *string `json:"user_id"`
}

type TaskForUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	StatusID    *string `json:"status_id"`
	Color       *string `json:"color"`
	UserID      *string `json:"user_id"`
}

func ListTasks(db *sql.DB) {
	sql, _, _ := goqu.From("task").ToSql()

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal("Error listing tasks")
	}

	defer rows.Close()

	for rows.Next() {
		task, _ := mapRowToTask(rows)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// CheckError(err)

		fmt.Println(task)
	}

	// return tasks
}

// impl chi.Render interface to render objects as JSON to the API consumer
func (t *Task) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (t *Task) insertTask(db *sql.DB) {
	title := t.Title
	sql, _, _ := goqu.From("tasks").Where(goqu.Ex{
		"title": []string{title},
	}).ToSql()

	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal("Error inserting task")
	}
}

func mapRowToTask(rows *sql.Rows) (Task, error) {
	var t Task
	err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.StatusID, &t.UserID)

	if err != nil {
		// if err.Is(err, sql.ErrNoRows) {
		// 	return Task{}, serrors.ErrNotFound
		// }
		// return Task{}, errors.Wrap(err, "failed to scan metadata")
	}

	return t, nil

}
