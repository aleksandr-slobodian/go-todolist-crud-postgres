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
	Users interface{
		Create(context.Context, *User) error
	}
	Todos interface{
		GetByID(context.Context, int64) (*Todo, error)
		Create(context.Context, *Todo) error
		Update(context.Context, *Todo) error
		Delete(context.Context, int64) error
		GetTodos(context.Context, TodosQueryParams) ([]*Todo, error)
	}

}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
		Todos: &TodoStore{db},
	}
}