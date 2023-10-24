package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "notesDB.sqlite")
	if err != nil {
		log.Fatalf("🔥 failed to connect to the database: %s", err.Error())
	}

	log.Println("🚀 Connected Successfully to the Database")

	return db
}

func MakerMigrations() error {
	db := GetConnection()

	stmt := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(64) NULL,
		description VARCHAR(255) NULL,
		completed BOOLEAN DEFAULT(FALSE),
		created_at TIMESTAMP DEFAULT DATETIME
	  );`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

/*
https://noties.io/blog/2019/08/19/sqlite-toggle-boolean/index.html
*/
