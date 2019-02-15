package pointerror

import (
	"fmt"
	"github.com/pkg/errors"
)

type Stringer interface {
	String() string
}
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func(w *Wallet) Deposit(amount Bitcoin){
	w.balance += Bitcoin(amount)
}

func(w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= Bitcoin(amount)
	return nil
}

func(w *Wallet) Balance() Bitcoin{
	return w.balance
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}
