package db

import (
	"context"
	"database/sql"
	"sort"
	"testing"
	"time"

	"github.com/amirrhkm/bank.be/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(owner string, balance int64, currency string) (Accounts, error) {
	arg := CreateAccountParams{
		Owner:    owner,
		Balance:  balance,
		Currency: currency,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	return account, err
}

func TestCreateAccount(t *testing.T) {
	user := createRandomUser(t)
	owner := user.Username
	balance := util.RandomMoney()
	currency := util.RandomCurrency()

	account, err := createTestAccount(owner, balance, currency)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, owner, account.Owner)
	require.Equal(t, balance, account.Balance)
	require.Equal(t, currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	user := createRandomUser(t)
	owner := user.Username
	balance := util.RandomMoney()
	currency := util.RandomCurrency()

	account, err := createTestAccount(owner, balance, currency)
	response, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, response)
	require.Equal(t, account.ID, response.ID)
	require.Equal(t, account.Owner, response.Owner)
	require.Equal(t, account.Balance, response.Balance)
	require.Equal(t, account.Currency, response.Currency)
	require.WithinDuration(t, account.CreatedAt, response.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	user := createRandomUser(t)
	owner := user.Username
	balance := util.RandomMoney()
	currency := util.RandomCurrency()

	account, err := createTestAccount(owner, balance, currency)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	response, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, response)
	require.Equal(t, account.ID, response.ID)
	require.Equal(t, arg.Balance, response.Balance)
	require.Equal(t, account.Currency, response.Currency)
	require.WithinDuration(t, account.CreatedAt, response.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	user := createRandomUser(t)
	owner := user.Username
	balance := util.RandomMoney()
	currency := util.RandomCurrency()

	account, err := createTestAccount(owner, balance, currency)

	err = testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	response, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, response)
}

func TestListAccounts(t *testing.T) {
	_, err := testDB.Exec("DELETE FROM entries")
	require.NoError(t, err)
	_, err = testDB.Exec("DELETE FROM transfers")
	require.NoError(t, err)
	_, err = testDB.Exec("DELETE FROM accounts")
	require.NoError(t, err)
	_, err = testDB.Exec("DELETE FROM users")
	require.NoError(t, err)

	var accountIDs []int64
	for i := 0; i < 10; i++ {
		user := createRandomUser(t)
		account, err := createTestAccount(user.Username, util.RandomMoney(), util.RandomCurrency())
		require.NoError(t, err)
		require.NotEmpty(t, account)
		accountIDs = append(accountIDs, account.ID)
	}

	sort.Slice(accountIDs, func(i, j int) bool {
		return accountIDs[i] < accountIDs[j]
	})

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for i, account := range accounts {
		require.NotEmpty(t, account)
		expectedID := accountIDs[i+5]
		require.Equal(t, expectedID, account.ID,
			"Account at position %d has ID %d, expected %d", i, account.ID, expectedID)
	}
}
