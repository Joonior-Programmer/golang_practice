package baccount

import (
	"errors"
	"fmt"
)

/*

	Start with Lowercase -> Private
	Start with Uppercase -> Public

	Method
		will have a receiver between func and funcName
		( ex. func (r Receiver) )

		Naming convention is lowercase of struct

*/

type bankAccount struct {
	owner   string
	balance int
}

// CreateBankAccount creates a new bank account
func CreateBankAccount(owner string) *bankAccount {
	newAccount := bankAccount{owner: owner, balance: 0}
	return &newAccount
}

// Deposit money into the account
func (b *bankAccount) Deposit(amount int) error {
	if amount < 0 {
		return errors.New("amount cannot be less than 0")
	}
	b.balance += amount
	fmt.Printf("Successfully deposited %d\n", amount)
	return nil
}

// Withdraw money from the account
func (b *bankAccount) Withdraw(amount int) error {
	if amount < 0 {
		return errors.New("amount cannot be less than 0")
	}
	if b.balance < amount {
		return errors.New("you cannot withdraw money more than your bank account has")
	}
	b.balance -= amount
	fmt.Printf("Successfully withdrew %d\n", amount)
	return nil
}

// Check Balance inside the account
func (b bankAccount) CheckBalance() int {
	return b.balance
}
