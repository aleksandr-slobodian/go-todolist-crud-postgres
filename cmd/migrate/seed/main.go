package main

import (
	"log"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/db"
	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/env"
	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/store"
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