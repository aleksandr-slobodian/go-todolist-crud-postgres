package db

import (
	"context"
	"database/sql"
	"log"
	"math/rand"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/cmd/internal/store"
	"github.com/go-faker/faker/v4"
)

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	todos := generateTodos(10)

	for _, todo := range todos {		
		if err := store.Todos.Create(ctx, todo); err != nil {
			log.Panic(err)
		}
	}
}

func generateTodos(num int) []*store.Todo {
	todos := make([]*store.Todo, num)
	for i := 0; i < num; i++ {
		todos[i] = &store.Todo{
			Item:     faker.Sentence(),
			Completed: rand.Intn(2) == 0,
		}
	}
	return todos
}
