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

	ok, err := tableExists(db, "patients")

	if err != nil {
		log.Fatal(err)
	}

	if !ok {
		err = createTables(db)
		if err != nil {
			log.Fatalf("error creating tables: %v", err)
		}
	}

	return db
}

func tableExists(db *sql.DB, table_name string) (bool, error) {
	query := "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?;"

	row := db.QueryRow(query, table_name)

	var count int
	err := row.Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func createTables(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, query := range tableSchema {
		_, err := tx.Exec(query)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
