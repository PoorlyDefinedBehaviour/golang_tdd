package wallet

import (
	"golang_tdd/src/wallet/amount"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Deposit(t *testing.T) {
	wallet := New()

	amount, err := amount.New(10)

	assert.Nil(t, err)

	wallet.Deposit(amount)

	assert.Equal(t, 10, wallet.Balance())
}
