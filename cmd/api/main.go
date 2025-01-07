package main

import (
	"log"

	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/db"
	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/env"
	"github.com/aleksandr-slobodian/go-todolist-crud-postgres/internal/store"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	version string
	addr string
	apiURL string
	db dbConfig
	env string
}

type dbConfig struct {
	addr string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime string
}

//	@title						Go Todo List CRUD Application
//	@description				This is a sample api for todo list application
//	@BasePath					/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description
func main() {
	
	cfg := config{
		version: env.GetString("VERSION", "0.0.0"),
		addr: env.GetString("ADDR", ":8181"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8181"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", ""),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}
	
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := 	&application{
		config: cfg,
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
	
}