package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/nuriansyah/rocket-ticket/db/migration"
)

func main() {
	db, err := sql.Open("sqlite3", "./basis-app.db")
	if err != nil {
		panic(err)
	}
	migration.Migrate(db)
}
