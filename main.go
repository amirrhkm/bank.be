package main

import (
	"database/sql"
	"log"

	"github.com/amirrhkm/bank.be/api"
	db "github.com/amirrhkm/bank.be/db/sqlc"
	"github.com/amirrhkm/bank.be/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
