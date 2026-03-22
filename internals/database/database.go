package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbName string = "smilelog"

func CreateDb() {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s.db", dbName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(tableSchema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Database %s.db created successfully!", dbName)
}
