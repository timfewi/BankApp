package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/timfewi/BankApp/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func createRandomAccountWithOwner(t *testing.T, owner string) Account {
	arg := CreateAccountParams{
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func deleteAccountsByTest(t *testing.T, owner string) {
	query := "DELETE FROM account WHERE owner = $1"
	_, err := testQueries.db.ExecContext(context.Background(), query, owner)
	require.NoError(t, err)
}

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	deleteAccountsByTest(t, arg.Owner)

}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
	deleteAccountsByTest(t, account1.Owner)
}

func TestGetsAccounts(t *testing.T) {
	uniqueOwner := util.RandomOwner()

	for i := 0; i < 10; i++ {
		createRandomAccountWithOwner(t, uniqueOwner)
	}

	arg := GetAccountsParams{
		Owner:  uniqueOwner,
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

	

	deleteAccountsByTest(t, uniqueOwner)
}


func TestGetsAccountsWithNoOffset(t *testing.T) {
    uniqueOwner := util.RandomOwner()
    for i := 0; i < 5; i++ {
        createRandomAccountWithOwner(t, uniqueOwner)
    }

    arg := GetAccountsParams{
        Owner:  uniqueOwner,
        Limit:  5,
        Offset: 0,
    }

    accounts, err := testQueries.GetAccounts(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, accounts, 5) 

    deleteAccountsByTest(t, uniqueOwner)
}


func TestGetsAccountsWithLimitExceeding(t *testing.T) {
    uniqueOwner := util.RandomOwner()
    totalAccounts := 3 
    for i := 0; i < totalAccounts; i++ {
        createRandomAccountWithOwner(t, uniqueOwner)
    }

    arg := GetAccountsParams{
        Owner:  uniqueOwner,
        Limit:  10, 
        Offset: 0,
    }

    accounts, err := testQueries.GetAccounts(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, accounts, totalAccounts) 

    deleteAccountsByTest(t, uniqueOwner)
}


func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountBalanceParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccountBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	deleteAccountsByTest(t, account1.Owner)
}
