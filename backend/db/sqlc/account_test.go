package db

import (
	"context"
	"testing"
	"time"

	"github.com/dhiegogoncalves/gofinance/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	category := createRandomCategory(t)

	arg := CreateAccountParams{
		UserID:      category.UserID,
		CategoryID:  category.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
		Value:       10.50,
		Date:        time.Now(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.CategoryID, account.CategoryID)
	require.Equal(t, arg.Title, account.Title)
	require.Equal(t, arg.Type, account.Type)
	require.Equal(t, arg.Description, account.Description)
	require.Equal(t, arg.Value, account.Value)
	require.NotEmpty(t, account.CreatedAt)
	require.NotEmpty(t, account.Date)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccountById(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccountById(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.CategoryID, account2.CategoryID)
	require.Equal(t, account1.Title, account2.Title)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Description, account2.Description)
	require.Equal(t, account1.Value, account2.Value)
	require.NotEmpty(t, account2.CreatedAt)
	require.NotEmpty(t, account2.Date)
}

func TestListAccounts(t *testing.T) {
	account := createRandomAccount(t)

	arg := GetAccountsParams{
		UserID:      account.UserID,
		Type:        account.Type,
		CategoryID:  account.CategoryID,
		Title:       account.Title,
		Description: account.Description,
		Date:        account.Date,
	}

	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, a := range accounts {
		require.Equal(t, a.ID, account.ID)
		require.Equal(t, a.UserID, arg.UserID)
		require.Equal(t, a.Title, arg.Title)
		require.Equal(t, a.Type, arg.Type)
		require.Equal(t, a.Description, arg.Description)
		require.Equal(t, a.Value, account.Value)
		require.NotEmpty(t, a.CategoryTitle)
		require.NotEmpty(t, a.CreatedAt)
		require.NotEmpty(t, a.Date)
	}
}

func TestAccountsReports(t *testing.T) {
	account := createRandomAccount(t)

	arg := GetAccountsReportsParams{
		UserID: account.UserID,
		Type:   account.Type,
	}

	sumValue, err := testQueries.GetAccountsReports(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sumValue)
}

func TestAccountsGraph(t *testing.T) {
	account := createRandomAccount(t)

	arg := GetAccountsGraphParams{
		UserID: account.UserID,
		Type:   account.Type,
	}

	graphValue, err := testQueries.GetAccountsGraph(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, graphValue)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:          account1.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(20),
		Value:       20.10,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.CategoryID, account2.CategoryID)
	require.Equal(t, arg.Title, account2.Title)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, arg.Description, account2.Description)
	require.Equal(t, arg.Value, account2.Value)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.NotEmpty(t, account2.Date)
}

func TestDeleteAccountById(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccountById(context.Background(), account.ID)
	require.NoError(t, err)

	_, err = testQueries.GetAccountById(context.Background(), account.ID)
	require.Error(t, err)
}
