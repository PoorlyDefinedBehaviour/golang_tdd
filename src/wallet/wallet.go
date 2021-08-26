package wallet

import (
	"sync"

	"github.com/pkg/errors"
	"golang_tdd/src/wallet/amount"
)

type balance = struct {
	value amount.T
	lock  sync.RWMutex
}

type T struct {
	balance balance
}

func New() T {
	value, err := amount.New(0)
	if err != nil {
		panic(err)
	}

	return T{
		balance: balance{
			value: value,
			lock:  sync.RWMutex{},
		},
	}
}

func (wallet *T) Deposit(value amount.T) {
	wallet.balance.lock.Lock()
	defer wallet.balance.lock.Unlock()

	wallet.balance.value = amount.Add(wallet.balance.value, value)
}

var ErrInsufficientFunds = errors.New("unsufficient funds")

func (wallet *T) Withdraw(value amount.T) error {
	wallet.balance.lock.Lock()
	defer wallet.balance.lock.Unlock()

	if amount.Int(value) > amount.Int(wallet.balance.value) {
		return errors.Wrapf(
			ErrInsufficientFunds,
			"tried to withdraw %d out of balance %d",
			amount.Int(value),
			amount.Int(wallet.balance.value),
		)
	}

	wallet.balance.value = amount.Sub(wallet.balance.value, value)

	return nil
}

func (wallet *T) Balance() amount.T {
	wallet.balance.lock.RLock()
	defer wallet.balance.lock.RUnlock()

	return wallet.balance.value
}
