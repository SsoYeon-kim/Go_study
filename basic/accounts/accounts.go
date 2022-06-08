package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner   string
	balance int
}

//NewAccount creates Accoun
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//method
// (a Account) : receiver
// a : Account type
// receiver는 struct 앞글자 소문자
// a :  Account의 복사본 (실제 account가 아님)
// *로 account를 복사하지 않고 Deposi method를 호출한 account를 사용함
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//balance of your account
func (a Account) Balance() int {
	return a.balance
}

var errnoMoney = errors.New("Can't withdraw")

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errnoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
