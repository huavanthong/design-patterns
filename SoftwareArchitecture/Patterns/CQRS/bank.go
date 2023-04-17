package main

import (
	"errors"
	"fmt"
)

// BankAccount represents a bank account with a balance.
type BankAccount struct {
	balance float64
}

/*
Experience 2:
Chúng ta có thể không cần define interface cho BankAccount struct bởi vì mục đích
interface lúc này là giúp ta có thể dễ quản lý và sử dụng ở các package khác.
*/
type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	GetBalance() float64
}

/*
Experience 1: Khi define các phương thức mà nhận input (a *BankAccount) thì khi
khởi tạo một instance BankAccount nó sẽ sử dụng được các phương thức đó
*/
// Deposit adds the specified amount to the account balance.
func (a *BankAccount) Deposit(amount float64) {
	a.balance += amount
}

// Withdraw subtracts the specified amount from the account balance, if sufficient funds are available.
func (a *BankAccount) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

// GetBalance returns the current balance of the account.
func (a *BankAccount) GetBalance() float64 {
	return a.balance
}

// BankAccountQuery is a class responsible for reading data from the BankAccount.
type BankAccountQuery struct {
	account *BankAccount
}

// GetBalance returns the current balance of the account.
func (q *BankAccountQuery) GetBalance() float64 {
	return q.account.GetBalance()
}

// BankAccountCommand is a class responsible for writing data to the BankAccount.
type BankAccountCommand struct {
	account *BankAccount
}

// Deposit adds the specified amount to the account balance.
func (c *BankAccountCommand) Deposit(amount float64) {
	c.account.Deposit(amount)
}

// Withdraw subtracts the specified amount from the account balance, if sufficient funds are available.
func (c *BankAccountCommand) Withdraw(amount float64) error {
	return c.account.Withdraw(amount)
}

// BankAccountView is a class responsible for displaying the BankAccount balance to the user.
type BankAccountView struct{}

// DisplayBalance prints the current balance of the account to the console.
func (v *BankAccountView) DisplayBalance(balance float64) {
	fmt.Printf("Current balance: %.2f\n", balance)
}

// BankAccountController is a class responsible for handling user input and coordinating between the view, query, and command objects.
type BankAccountController struct {
	view    *BankAccountView
	query   *BankAccountQuery
	command *BankAccountCommand
}

// NewBankAccountController creates a new BankAccountController with the specified BankAccount.
func NewBankAccountController(account *BankAccount) *BankAccountController {
	return &BankAccountController{
		view:    &BankAccountView{},
		query:   &BankAccountQuery{account},
		command: &BankAccountCommand{account},
	}
}

// DepositHandler handles a deposit request from the user.
func (c *BankAccountController) DepositHandler(amount float64) {
	c.command.Deposit(amount)
	c.view.DisplayBalance(c.query.GetBalance())
}

// WithdrawHandler handles a withdraw request from the user.
func (c *BankAccountController) WithdrawHandler(amount float64) {
	err := c.command.Withdraw(amount)
	if err != nil {
		fmt.Println(err)
	} else {
		c.view.DisplayBalance(c.query.GetBalance())
	}
}

func main() {
	account := &BankAccount{}
	fmt.Printf("Init account with balance: %.2f\n", account.GetBalance())

	controller := NewBankAccountController(account)

	controller.DepositHandler(100)
	controller.WithdrawHandler(50)

	balance := controller.query.GetBalance()
	fmt.Printf("Current balance: %.2f\n", balance)
}
