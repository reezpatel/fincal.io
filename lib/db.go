package lib

import "database/sql"

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/roger-king/go-getenv"
)

var (
	DB *sql.DB
)

func initDB() {
	dbURL := getenv.EnvOrDefault("PG_LOCAL_DB_URL", getenv.String("localhost"))

	db, err := sql.Open("pgx", dbURL)

	if err != nil {
		panic(err)
	}

	DB = db
}
