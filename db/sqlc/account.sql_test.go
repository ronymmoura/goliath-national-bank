package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ronymmoura/goliath-national-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		UserID:   user.ID,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	account, err := testStore.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAccount.ID, account.ID)
	require.Equal(t, createdAccount.UserID, account.UserID)
	require.Equal(t, createdAccount.Balance, account.Balance)
	require.Equal(t, createdAccount.Currency, account.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		UserID: lastAccount.UserID,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.UserID, account.UserID)
	}
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: util.RandomMoney(),
	}

	account, err := testStore.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAccount.ID, account.ID)
	require.Equal(t, createdAccount.UserID, account.UserID)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, createdAccount.Currency, account.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestSelectAccountForUpdate(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: util.RandomMoney(),
	}

	account, err := testStore.GetAccountForUpdate(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	account, err = testStore.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAccount.ID, account.ID)
	require.Equal(t, createdAccount.UserID, account.UserID)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, createdAccount.Currency, account.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	err := testStore.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	account, err := testStore.GetAccount(context.Background(), createdAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, account)
}
