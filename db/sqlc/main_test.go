package db

import (
	"database/sql"
	"github.com/kartheekakkur/service/utils"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../.")

	if err != nil {
		log.Fatal("Cannot load the config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the Database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
