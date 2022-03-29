package main

import "fmt"

type bankAccount struct {
	ownerName            string
	identificationNumber uint64
	branch               string
	balance              int64
}

type BankAccount interface {
	WithDraw(amt uint64)
	Deposit(amt uint64)
	GetBalance() uint64
}

type BankAccountBuilder interface {
	WithOwnerName(name string) BankAccountBuilder
	WithOwnerIdentity(identificationNumber uint64) BankAccountBuilder
	AtBranch(branch string) BankAccountBuilder
	OpeningBalance(balance uint64) BankAccountBuilder
	Build() BankAccount
}

func (acc *bankAccount) WithDraw(amt uint64) {

	if acc.balance > 0 {
		temp := acc.balance - int64(amt)
		if temp < 0 || acc.balance < 50000 {
			fmt.Println("Can't get money greater your balance")
		} else {
			acc.balance = temp
			fmt.Println("With draw success")
		}

	} else {
		fmt.Println("No enough money")
	}
}

func (acc *bankAccount) Deposit(amt uint64) {

	if acc.balance > 0 {
		acc.balance = acc.balance + int64(amt)
	} else {
		fmt.Println("No enough money")
	}
}

func (acc *bankAccount) GetBalance() uint64 {
	return 0
}

func (acc *bankAccount) WithOwnerName(name string) BankAccountBuilder {
	acc.ownerName = name
	return acc
}

func (acc *bankAccount) WithOwnerIdentity(identificationNumber uint64) BankAccountBuilder {
	acc.identificationNumber = identificationNumber
	return acc
}

func (acc *bankAccount) AtBranch(branch string) BankAccountBuilder {
	acc.branch = branch
	return acc
}

func (acc *bankAccount) OpeningBalance(balance uint64) BankAccountBuilder {
	acc.balance = int64(balance)
	return acc
}

func (acc *bankAccount) Build() BankAccount {
	return acc
}

func NewBankAccountBuilder() BankAccountBuilder {
	return &bankAccount{}
}

func main() {
	account := NewBankAccountBuilder().
		WithOwnerName("Tuan").
		WithOwnerIdentity(123456789).
		AtBranch("Sai Gon").
		OpeningBalance(1000).Build()

	account.Deposit(10000)
	account.WithDraw(50000)

	fmt.Print(account)
}
