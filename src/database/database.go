package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(10)

	return db
}
