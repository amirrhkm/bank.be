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

var testQueries *Queries
var testDB *sql.DB

/*
 * Unit Tests Prerequisite:
 * Ensure database `bank-test` is created and migrated
 */
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.TestDBSource)
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
