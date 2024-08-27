package db

import (
	"context"
	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func getAccounts(t *testing.T) []Account {
	arg := ListAccountsParams{
		Limit:  2,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	return accounts
}

func createTransferTest(t *testing.T) Transfer {
	account := getAccounts(t)
	account1 := account[0]
	account2 := account[1]
	require.NotEmpty(t, account1)
	require.NotEmpty(t, account2)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	createTransferTest(t)
}

func TestQueries_GetTransfer(t *testing.T) {
	transferRecord := createTransferTest(t)
	require.NotEmpty(t, transferRecord)
	transfer, err := testQueries.GetTransfer(context.Background(), transferRecord.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
}

func TestQueries_ListTransfers(t *testing.T) {
	account := getAccounts(t)
	account1 := account[0]
	account2 := account[1]
	require.NotEmpty(t, account1)
	require.NotEmpty(t, account2)

	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         1,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)
}
