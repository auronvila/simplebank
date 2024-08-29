package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/simplebank/api"
	db "github.com/simplebank/db/sqlc"
	"log"
)

func main() {
	const (
		dbDriver      = "postgres"
		dbSource      = "postgres://postgres:1@localhost:5432/simple-bank?sslmode=disable"
		serverAddress = "0.0.0.0:3002"
	)

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
