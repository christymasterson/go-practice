package bankaccount

import (
	"fmt"
)

var balance int = 0

func deposit(amount int) int {
	balance += amount
	return balance
}

func withdraw(amount int) int {
	balance -= amount
	return balance
}

func showbalance() int {
	fmt.Println(balance)
	return balance
}
