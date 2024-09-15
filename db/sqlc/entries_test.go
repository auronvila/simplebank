package db

import (
	"context"
	"fmt"
	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func getSingleAccount(t *testing.T) Account {
	hashedPass, _ := util.HashPassword(util.RandomString(5))
	user, _ := testQueries.CreateUser(context.Background(), CreateUserParams{
		Username:       util.RandomString(5),
		HashedPassword: hashedPass,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	})
	_, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    user.Username,
		Balance:  0,
		Currency: util.RandomCurrency(),
	})
	require.NoError(t, err)

	if err != nil {
		return Account{}
	}

	arg := ListAccountsParams{
		Owner:  user.Username,
		Limit:  1,
		Offset: 0,
	}
	fmt.Println("ARGSS", arg)
	fmt.Println("Username", user.Username)
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	fmt.Println("Accounts", accounts)
	require.NoError(t, err)
	//require.NotEmpty(t, accounts)
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
