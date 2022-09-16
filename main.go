package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-aws-micro/api"
	db "go-aws-micro/db/sqlc"
	"go-aws-micro/util"

	"log"
)

/*const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)*/

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
