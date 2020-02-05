package bankaccount

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	balance = 0
	amount := 10
	deposit(amount)
	if balance != amount {
		t.Errorf("Expected %v in bank account, got %v", amount, balance)
	}
}

func TestWithdraw(t *testing.T) {
	balance = 10
	amount := 5
	withdraw(amount)
	if balance != 5 {
		t.Errorf("Expected 5 in bank account, got %v", balance)
	}
}

func TestBalance(t *testing.T) {
	balance = 10
	if showbalance() != 10 {
		t.Errorf("Expected to print balance 10")
	}
}
