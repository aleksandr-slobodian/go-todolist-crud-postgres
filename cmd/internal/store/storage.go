package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Todos interface{
		GetByID(context.Context, int64) (*Todo, error)
		Create(context.Context, *Todo) error
		Update(context.Context, *Todo) error
		Delete(context.Context, int64) error
		GetTodos(context.Context) ([]*Todo, error)
	}

}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Todos: &TodoStore{db},
	}
}