package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rinuccia/transaction-system/config"
	"os"
)

var schema = `
	CREATE TABLE IF NOT EXISTS users (
                         id serial PRIMARY KEY,
                         first_name varchar NOT NULL,
                         last_name varchar NOT NULL,
                         email varchar UNIQUE NOT NULL,
                         balance bigint NOT NULL
	);`

func NewPostgresClient(cfg *config.Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.Username, os.Getenv("DB_PASSWORD"), cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
