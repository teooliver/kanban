package user

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/doug-martin/goqu/v9"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) ListAllUsers(ctx context.Context) ([]User, error) {
	sql, _, err := goqu.From("app_user").Select(allColumns...).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("error generating list all user query: %w", err)
	}

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, fmt.Errorf("error executing list all user query: %w", err)
	}

	defer rows.Close()

	var result []User
	for rows.Next() {
		user, err := mapRowToUser(rows)
		if err != nil {
			return nil, err
		}
		slog.Info("LIST SQL RESULT ===> %+v\n", "result", user)

		result = append(result, user)
	}

	return result, nil
}

func (r *PostgresRepository) CreateUser(ctx context.Context, user UserForCreate) (err error) {
	insertSQL, args, err := goqu.Insert("app_user").Rows(UserForCreate{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating create user query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing create user query: %w", err)
	}

	slog.Info("CREATE RESULT", result)
	return nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, userID string) (err error) {
	insertSQL, args, err := goqu.Delete("app_user").Where(goqu.Ex{"id": userID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating delete user query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, insertSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing delete user query: %w", err)
	}

	slog.Info("DELETED USER ID", result)
	return nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, userID string, user UserForUpdate) (err error) {
	updateSQL, args, err := goqu.Update("app_user").Set(user).Where(goqu.Ex{"id": userID}).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("error generating update user query: %w", err)
	}

	result, err := r.db.ExecContext(ctx, updateSQL, args...)
	if err != nil {
		return fmt.Errorf("error executing update user query: %w", err)
	}

	slog.Info("UPDATED ID", result)
	return nil
}
