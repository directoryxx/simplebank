package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createAccountTest(t *testing.T) Accounts {
	arg := CreateAccountParams{
		Owner:    "angga",
		Balance:  1000,
		Currency: "IDR",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createAccountTest(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createAccountTest(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	require.NotEmpty(t, account2)

	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestListAccount(t *testing.T) {
	createAccountTest(t)
	accountList := ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}
	accountListData, err := testQueries.ListAccounts(context.Background(), accountList)
	require.NoError(t, err)

	require.NotEmpty(t, accountListData)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createAccountTest(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: 10000,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, account2)

	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, arg.ID, account2.ID)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createAccountTest(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	getAccount, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.Empty(t, getAccount)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
