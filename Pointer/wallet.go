package Pointer

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

// string型にフォーマットされたとき、自動的にBTCが末尾につくようになる
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// こっちはポインタにする必要はないが、慣例的に一貫性を保つため書くようにする
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
