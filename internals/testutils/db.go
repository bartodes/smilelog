package testutils

import (
	"database/sql"
	"testing"

	"github.com/bartodes/smilelog/internals/database"
)

func SetupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	database.ExecSchema(db)

	return db
}
