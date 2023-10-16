package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // because we don't directly call this, we need to add _ ahead of it so that it is included
	"github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

// Test main is the entry point for all golang tests for a specific package
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot read environment config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run()) // This tells unit tests that they can start running, and will return an exit code which should go to os.
}
