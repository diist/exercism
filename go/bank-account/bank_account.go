// Package account contains logic for the bank account Exercism exercise
package account

import "sync"

// Account is a bank account with a balance, an open flag and a mutex for concurrency
type Account struct {
	balance int64
	open    bool
	sync.Mutex
}

// Open opens a bank account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, open: true}
}

// Close closes a bank account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}

	payout = a.balance
	a.balance = 0
	a.open = false
	return payout, true
}

// Balance returns the balance of a bank account
func (a *Account) Balance() (balance int64, ok bool) {
	return a.balance, a.open
}

// Deposit deposits money into a bank account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return a.balance, false
	}

	newBalance = a.balance + amount
	if newBalance < 0 {
		return a.balance, false
	}

	a.balance = newBalance
	return a.balance, true
}
