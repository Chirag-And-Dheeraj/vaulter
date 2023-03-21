package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	log.Println("Initializing SQLite database...")
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	table, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users
	(
		id TEXT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		email TEXT,
		pin TEXT,
		profile_photo BLOB,
		profile_creation_time TEXT,
		active TINYINT
	)`)


	if err != nil {
		log.Fatal(err)
	}
	defer table.Close()
	table.Exec()

	log.Println("Database initialized.")
	return db
}