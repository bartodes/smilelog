package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB_NAME string = "smilelog"

func InitDb(db_path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", db_path)

	if err != nil {
		return db, fmt.Errorf("error opening database: %w", err)
	}

	ok, err := tableExists(db, "patients")

	if err != nil {
		return db, err
	}

	if !ok {
		err = createTables(db)
		if err != nil {
			return db, fmt.Errorf("error creating tables: %w", err)
		}
	}

	return db, nil
}

func GetDbPath() string {
	base_path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/%s.db", base_path, DB_NAME)
}

func tableExists(db *sql.DB, table_name string) (bool, error) {
	query := "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?"

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
