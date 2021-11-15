package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	open    bool
	lock    *sync.Mutex
}

func Open(amount int64) *Account {
	var a *Account
	if amount < 0 {
		return a
	}

	return &Account{
		balance: amount,
		open:    true,
		lock:    &sync.Mutex{},
	}
}

func (a *Account) Close() (payout int64, ok bool) {

	a.lock.Lock()
	defer a.lock.Unlock()

	if !a.open {
		return 0, false
	}

	if a.balance >= 0 {
		payout = a.balance
		a.balance = 0
		a.open = false
		ok = true
	} else {
		payout = 0
		ok = false
	}

	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {

	a.lock.Lock()
	defer a.lock.Unlock()

	if !a.open {
		return 0, false
	}

	ok = true
	newBalance = a.balance
	newBalance += amount
	if newBalance < 0 {
		ok = false
	} else {
		a.balance = newBalance
	}

	return
}

func (a *Account) Balance() (balance int64, ok bool) {

	a.lock.Lock()
	defer a.lock.Unlock()

	if !a.open {
		return 0, false
	}

	balance = a.balance
	ok = true

	if balance < 0 {
		ok = false
	}

	return
}
