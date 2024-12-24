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

	users := generateUsers(3)

	for _, user := range users {		
		if err := store.Users.Create(ctx, user); err != nil {
			log.Panic(err)
		}
	}

	todos := generateTodos(10, users)

	for _, todo := range todos {		
		if err := store.Todos.Create(ctx, todo); err != nil {
			log.Panic(err)
		}
	}
}

func generateTodos(num int, users []*store.User) []*store.Todo {
	todos := make([]*store.Todo, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		todos[i] = &store.Todo{
			UserID: user.ID,
			Item:     faker.Sentence(),
			Completed: rand.Intn(2) == 0,
		}
	}
	return todos
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: faker.Username(),
			Email: faker.Email(),
			Password: faker.Password(),
		}
	}
	return users
}
