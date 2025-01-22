package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amirrhkm/bank.be/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entries {
	account := createRandomAccount(t)
	amount := util.RandomMoney()
	if util.RandomBool() {
		amount = -amount
	}

	arg := CreateEntryParams{
		AccountID: sql.NullInt64{Int64: account.ID, Valid: true},
		Amount:    amount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		amount := util.RandomMoney()
		if util.RandomBool() {
			amount = -amount
		}

		arg := CreateEntryParams{
			AccountID: sql.NullInt64{Int64: account.ID, Valid: true},
			Amount:    amount,
		}

		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg := ListEntriesParams{
		AccountID: sql.NullInt64{Int64: account.ID, Valid: true},
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID.Int64)
	}
}
