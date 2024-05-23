package deps

import (
	"database/sql"
	"log"

	"github.com/teooliver/kanban/internal/config"
)

type Infra struct {
	Postgres *sql.DB
}

func InitInfra(cfg *config.Config) (*Infra, error) {
	// TODO: handle err
	pgClient, _ := initPostgres(&cfg.Postgres)

	return &Infra{
		Postgres: pgClient,
	}, nil
}

func initPostgres(cfg *config.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		// TODO: Better error handling
		log.Fatal("Error connecting to db")
		return db, err
	}
	log.Println("Database connection established")

	defer db.Close()

	return db, nil
}