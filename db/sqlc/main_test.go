package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/techschool/simplebank/util"
)

var testStore Store

// Test main is the entry point for all golang tests for a specific package
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot read environment config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run()) // This tells unit tests that they can start running, and will return an exit code which should go to os.
}
