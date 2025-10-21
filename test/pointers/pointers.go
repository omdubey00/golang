package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin // This is private outside this package as this is lowercase.
}

func (w *Wallet) Deposit(amount Bitcoin) { // these are called struct pointers and they are automatically dereferenced.
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficinetFunds = errors.New("insufficient funds to withdraw")

func (w *Wallet) Withdraw(amount Bitcoin) (err error) {
	if amount > w.Balance() {
		return ErrInsufficinetFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
