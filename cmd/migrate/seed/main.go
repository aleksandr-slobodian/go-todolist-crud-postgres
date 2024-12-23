package main

import (
	"log"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/cmd/internal/db"
	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/cmd/internal/env"
	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/cmd/internal/store"
)


func main() {
	var addr = env.GetString("DB_ADDR", "")

	conn, err := db.New(
		addr,
		3,
		3,
		"5m",
	)
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}