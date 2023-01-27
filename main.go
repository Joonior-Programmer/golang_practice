package main

import (
	"fmt"
	"golang_practice/baccount" // import local package
)

func main() {
	/*
		bankAccount struct is private.
		Therefore, I created a constructor function "CreateBankAccount" publicly.
	*/

	account1 := baccount.CreateBankAccount("Joon")

	err := account1.Deposit(-1)
	if err != nil {
		fmt.Println(err)
	}

	err = account1.Deposit(10)
	if err != nil {
		fmt.Println(err)
	}

	err = account1.Withdraw(-1)
	if err != nil {
		fmt.Println(err)
	}

	err = account1.Withdraw(11)
	if err != nil {
		fmt.Println(err)
	}

	err = account1.Withdraw(5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account1)
}
