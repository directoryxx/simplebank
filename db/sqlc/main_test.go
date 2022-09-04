package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	DSN      = "postgresql://postgres:YadasdasavC@127.0.0.1:30200/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, DSN)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
