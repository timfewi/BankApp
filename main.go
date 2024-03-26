package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/timfewi/BankApp/api"
	db "github.com/timfewi/BankApp/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://konta:1683@localhost:5433/bankApp?sslmode=disable"
	serverAddress = "0.0.0.0:8181"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatal("cannot ping db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
