package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(testDB)
	accoun1 := createRandomAccount(t)
	accoun2 := createRandomAccount(t)

	n := 2
	amount := int64(10)

	// run n concurrent transactions
	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: accoun1.ID,
				ToAccountID:   accoun2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	done := make(map[int]bool)

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfers
		require.NotEmpty(t, result.Transfer)
		require.Equal(t, accoun1.ID, result.Transfer.FromAccountID)
		require.Equal(t, accoun2.ID, result.Transfer.ToAccountID)
		require.Equal(t, amount, result.Transfer.Amount)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)

		// check FromEntry
		require.NotEmpty(t, result.FromEntry)
		require.Equal(t, accoun1.ID, result.FromEntry.AccountID)
		require.Equal(t, -amount, result.FromEntry.Amount)
		require.NotZero(t, result.FromEntry.ID)
		require.NotZero(t, result.FromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)

		// check ToEntry
		require.NotEmpty(t, result.ToEntry)
		require.Equal(t, accoun2.ID, result.ToEntry.AccountID)
		require.Equal(t, amount, result.ToEntry.Amount)
		require.NotZero(t, result.ToEntry.ID)
		require.NotZero(t, result.ToEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)

		// check FromAccount
		require.NotEmpty(t, result.FromAccount)
		require.Equal(t, accoun1.ID, result.FromAccount.ID)

		// check ToAccount
		require.NotEmpty(t, result.ToAccount)
		require.Equal(t, accoun2.ID, result.ToAccount.ID)

		// check balance
		fmt.Println("after tx: ", result.FromAccount.Balance, result.ToAccount.Balance)
		diff1 := accoun1.Balance - result.FromAccount.Balance
		diff2 := result.ToAccount.Balance - accoun2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k > 1 && k < n)
		require.Contains(t, done, k)
		done[k] = true

	}

	// check final updated balances
	updateAccount1, err := store.GetAccount(context.Background(), accoun1.ID)
	require.NoError(t, err)
	updateAccount2, err := store.GetAccount(context.Background(), accoun2.ID)
	require.NoError(t, err)

	fmt.Println("after: ", updateAccount1.Balance, updateAccount2.Balance)

	require.Equal(t, accoun1.Balance-int64(n)*amount, updateAccount1)
	require.Equal(t, accoun2.Balance-int64(n)*amount, updateAccount2)

}
