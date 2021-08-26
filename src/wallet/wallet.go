package wallet

import (
	"golang_tdd/src/wallet/amount"
	"sync"
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

func (wallet *T) Balance() amount.T {
	wallet.balance.lock.RLock()
	defer wallet.balance.lock.RUnlock()

	return wallet.balance.value
}
