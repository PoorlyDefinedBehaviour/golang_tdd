package wallet

import (
	"golang_tdd/src/wallet/amount"
	"testing"

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
