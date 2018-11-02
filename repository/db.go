package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	)

func InitDB(filepath string) *sql.DB {
	db, error := sql.Open("sqlite3", filepath)

	if error != nil {
		panic(error)
	}

	if db == nil {
		panic("Database Nil")
	}

	// 1. migrate db
	migrate(db)

	return db
}

func migrate(db *sql.DB) {
	sql := `
			CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			done INTEGER NOT NULL
		);
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
