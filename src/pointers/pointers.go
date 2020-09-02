package pointers

import "fmt"

// Bitcoin - type
type Bitcoin int

type Stringer interface {
	String() string
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit - deposit
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw - withdraw
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance - balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
