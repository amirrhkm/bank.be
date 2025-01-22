package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/amirrhkm/bank.be/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/bank-test?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

/*
 * Unit Tests Prerequisite:
 * Ensure database `bank-test` is created and migrated
 */
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Failed connecting to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

func createRandomAccount(t *testing.T) Accounts {
	owner := util.RandomOwner()
	balance := util.RandomMoney()
	currency := util.RandomCurrency()

	account, err := createTestAccount(owner, balance, currency)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}
