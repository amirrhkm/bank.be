package db

import (
	"context"
	"testing"
	"time"

	"github.com/amirrhkm/bank.be/util"
	"github.com/stretchr/testify/require"
)

func createTestEntry(accountID int64, amount int64) (Entries, error) {
	arg := CreateEntryParams{
		AccountID: accountID,
		Amount:    amount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	return entry, err
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	amount := util.RandomOperation(util.RandomMoney())

	entry, err := createTestEntry(account.ID, amount)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	amount := util.RandomOperation(util.RandomMoney())

	entry, err := createTestEntry(account.ID, amount)
	response, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, entry.ID, response.ID)
	require.Equal(t, entry.AccountID, response.AccountID)
	require.Equal(t, entry.Amount, response.Amount)
	require.WithinDuration(t, entry.CreatedAt, response.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		amount := util.RandomOperation(util.RandomMoney())
		entry, err := createTestEntry(account.ID, amount)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}
}
