package baccount

/*
	Function or Struct
		Start with Lowercase -> Private
		Start with Uppercase -> Public
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
