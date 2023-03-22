package database

import (
	"log"

	M "server/database/models"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	log.Println("Initializing SQLite database...")
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&M.User{})

	log.Println("Database initialized.")
	return db
}
