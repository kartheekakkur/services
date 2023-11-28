package main

import (
	"database/sql"
	"log"

	"github.com/kartheekakkur/service/api"
	db "github.com/kartheekakkur/service/db/sqlc"
	"github.com/kartheekakkur/service/utils"
	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load the config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
