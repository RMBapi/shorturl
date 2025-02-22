package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {

	DB, err = sql.Open("sqlite3", "shorturl.db")

	if err != nil {
		panic("Couldn't connect database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {

	createTodoTable := `
	CREATE TABLE IF NOT EXISTS urllist (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    uniqueKey TEXT NOT NULL,
	shorturl TEXT NOT NULL
    )
	`
	_, err := DB.Exec(createTodoTable)

	if err != nil {
		panic("Couldn't connect event table")
	}

}
