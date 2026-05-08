package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB_NAME string = "smilelog"

func InitDB() *sql.DB {
	basePath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dbPath := fmt.Sprintf("%s/%s.db", basePath, DB_NAME)
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	ExecSchema(db)

	return db
}

func ExecSchema(db *sql.DB) {
	var count int

	err := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='patients';").Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	for _, query := range tableSchema {
		_, err := tx.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
