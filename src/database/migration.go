package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate() {
	db := GetDatabase()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)

	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://src/database/migrations/",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Steps(5)
}
