package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, t transaction) error {
	switch t.transactionType {
	case "deposit":
		c.balance += t.amount
	case "withdrawal":
		if t.amount > c.balance {
			return errors.New("insufficient funds")
		}
		c.balance -= t.amount
	default:
		return errors.New("unknown transaction type")
	}
	return nil
}

func main() {
	runTestCases()
}
// algorithm
// ------------------------------------
// 1. Input should be a pointer to customer and a transaction and output should be error
// 2. Depending on the transaction type:
//   - Withdrawal: subtract amount from customer's balance
// 	 - Deposite: add amount  to customer's balance
//   - If customer doesn't have enough money return error "insufficient funds"
//   - Otherwise return error "unknown transaction type"