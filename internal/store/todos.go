package store

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)


type Todo struct {
	ID        int64     `json:"id"`
	Item     	string    `json:"item"`
	Completed bool   		`json:"completed"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	UserID    int64     `json:"user_id"`
}

type TodosQueryParams struct {
	Limit int `form:"limit" binding:"gte=1,lte=100"`
	Offset int `form:"offset" binding:"omitempty,gte=0"`
	Order string `form:"order" binding:"omitempty,oneof=asc desc"`
	Search string `form:"search" binding:"omitempty,max=100"`
}

type 	TodoStore struct {
	db *sql.DB
}

func (s *TodoStore) Create(ctx context.Context, todo *Todo) error {
	query := `
		INSERT INTO todos (item, completed, user_id)
		VALUES ($1, $2, $3) RETURNING id, item, completed, created_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		todo.Item,
		todo.Completed,
		todo.UserID,
	).Scan(
		&todo.ID,
		&todo.Item,
		&todo.Completed,
		&todo.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoStore) Update(ctx context.Context, todo *Todo) error {
	query := `
		UPDATE todos
		SET item = $1, completed = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING item, completed, updated_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		todo.Item,
		todo.Completed,
		todo.ID,
	).Scan(
		&todo.Item, 
		&todo.Completed, 
		&todo.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}
	return nil
}

func (s *TodoStore) Delete(ctx context.Context, todoID int64) error {
	query := `DELETE FROM todos WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, todoID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *TodoStore) GetByID(ctx context.Context, id int64) (*Todo, error) {
	query := `
		SELECT id, item, completed, created_at, updated_at, user_id
		FROM todos
		WHERE id = $1
	`
	todo := Todo{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&todo.ID,
		&todo.Item,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.UserID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &todo, nil
}

func (s *TodoStore) GetTodos(ctx context.Context, qparams TodosQueryParams) ([]*Todo, error) {
	query := `
		SELECT id, item, completed, created_at, updated_at, user_id
		FROM todos
		WHERE item ILIKE '%' || $3 || '%'
		ORDER BY created_at ` + sortOrder(qparams.Order) + `, item ASC
		LIMIT $1
		OFFSET $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, qparams.Limit, qparams.Offset, qparams.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]*Todo, 0)

	for rows.Next() {
		todo := &Todo{}
		if err := rows.Scan(
			&todo.ID,
			&todo.Item,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.UserID,
		); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
