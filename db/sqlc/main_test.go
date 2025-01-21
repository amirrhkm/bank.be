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
	dbSource = "postgresql://root:password@localhost:5432/bank-test?sslmode=disable"
)

var testQueries *Queries

/*
 * Unit Tests Prerequisite:
 * Ensure database `bank-test` is created and migrated
 */
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Failed connecting to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
