package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/simplebank/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDb *sql.DB
var err error

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot load configs in test: ", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
