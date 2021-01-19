package account

import "sync"

// Account Simulate a bank account supporting opening/closing, withdrawals, and deposits of money.
type Account struct {
	sync.Mutex
	balance  int64
	isClosed bool
}

// Open create a new bank account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

// Close will close an existing bank account
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.isClosed {
		return 0, false
	}
	a.isClosed = true
	payout = a.balance
	a.balance = 0
	return payout, true
}

// Balance return the balance of a bank account
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.isClosed {
		return 0, false
	}
	return a.balance, true
}

// Deposit withdrawals and deposits of money.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.isClosed || a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}
