package wallet

import (
	"golang_tdd/src/wallet/amount"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_Deposit(t *testing.T) {
	wallet := New()

	deposit, err := amount.New(10)

	assert.Nil(t, err)

	expected, err := amount.New(10)

	assert.Nil(t, err)

	wallet.Deposit(deposit)

	assert.Equal(t, expected, wallet.Balance())
}

func Test_Withdraw(t *testing.T) {
	t.Parallel()

	t.Run("withdrawing", func(t *testing.T) {
		t.Parallel()

		wallet := New()

		initialValue, err := amount.New(10)

		assert.Nil(t, err)

		wallet.Deposit(initialValue)

		amountToWithdraw, err := amount.New(5)

		assert.Nil(t, err)

		expected, err := amount.New(5)

		assert.Nil(t, err)

		err = wallet.Withdraw(amountToWithdraw)

		assert.Nil(t, err)

		actual := wallet.Balance()

		assert.Equal(t, expected, actual)
	})

	t.Run("can't withdraw more than the current balance", func(t *testing.T) {
		t.Parallel()

		wallet := New()

		amountToWithdraw, err := amount.New(1)

		assert.Nil(t, err)

		err = wallet.Withdraw(amountToWithdraw)

		assert.True(t, errors.Is(err, ErrInsufficientFunds))
	})
}
