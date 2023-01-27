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
	fmt.Println(account1)
}
