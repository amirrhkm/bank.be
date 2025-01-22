package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/amirrhkm/bank.be/util"
	"github.com/stretchr/testify/require"
)

func createTestTransfer(fromAccountID int64, toAccountID int64, amount int64) (Transfers, error) {
	arg := CreateTransferParams{
		FromAccountID: sql.NullInt64{Int64: fromAccountID, Valid: true},
		ToAccountID:   sql.NullInt64{Int64: toAccountID, Valid: true},
		Amount:        amount,
	}

	response, err := testQueries.CreateTransfer(context.Background(), arg)

	return response, err
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := util.RandomMoney()

	transfer, err := createTestTransfer(account1.ID, account2.ID, amount)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, account1.ID, transfer.FromAccountID.Int64)
	require.Equal(t, account2.ID, transfer.ToAccountID.Int64)
	require.Equal(t, amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := util.RandomMoney()

	transfer, err := createTestTransfer(account1.ID, account2.ID, amount)
	response, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, response)

	require.Equal(t, transfer.ID, response.ID)
	require.Equal(t, transfer.FromAccountID, response.FromAccountID)
	require.Equal(t, transfer.ToAccountID, response.ToAccountID)
	require.Equal(t, transfer.Amount, response.Amount)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createTestTransfer(account1.ID, account2.ID, util.RandomMoney())
	}

	arg := ListTransfersParams{
		FromAccountID: sql.NullInt64{Int64: account1.ID, Valid: true},
		ToAccountID:   sql.NullInt64{Int64: account2.ID, Valid: true},
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Equal(t, len(transfers), 5)
}
