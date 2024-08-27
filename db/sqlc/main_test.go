package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:1@localhost:5432/simple-bank?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB
var err error

func TestMain(m *testing.M) {
	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
