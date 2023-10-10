package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // because we don't directly call this, we need to add _ ahead of it so that it is included
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:qwerty36@localhost:5432/simple_bank?sslmode=disable"
)

// Test main is the entry point for all golang tests for a specific package
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannon connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run()) // This tells unit tests that they can start running, and will return an exit code which should go to os.
}
