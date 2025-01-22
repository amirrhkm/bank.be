package db

import (
	"context"
	"database/sql"
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
	owner := util.RandomOwner()
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
	owner := util.RandomOwner()
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
	owner := util.RandomOwner()
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
	owner := util.RandomOwner()
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
	for i := 0; i < 10; i++ {
		createTestAccount(util.RandomOwner(), util.RandomMoney(), util.RandomCurrency())
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
