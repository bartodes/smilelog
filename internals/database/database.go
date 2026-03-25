package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB_NAME string = "smilelog"

func InitDb() {
	db_path := getDbPath()
	db := openDb(db_path)
	defer db.Close()

	ok := tableExists(db, "patients")

	if !ok {
		execDbSchema(db)
	}
}

func getDbPath() string {
	base_path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/%s.db", base_path, DB_NAME)
}

func openDb(db_path string) *sql.DB {
	db, err := sql.Open("sqlite3", db_path)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func tableExists(db *sql.DB, table_name string) bool {
	query := "SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?"

	row := db.QueryRow(query, table_name)

	var count int
	err := row.Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func execDbSchema(db *sql.DB) {
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

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
