package db

import (
	"context"
	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func getSingleAccount(t *testing.T) Account {
	arg := ListAccountsParams{
		Limit:  1,
		Offset: 1,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	return accounts[0]
}

func createEntryTest(t *testing.T) Entry {
	account := getSingleAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	createEntryTest(t)
}

func TestQueries_GetEntry(t *testing.T) {
	entry := createEntryTest(t)
	entry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
}

func TestQueries_ListEntries(t *testing.T) {
	entry := createEntryTest(t)
	arg := ListEntriesParams{
		AccountID: entry.AccountID,
		Limit:     1,
		Offset:    0,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
}
