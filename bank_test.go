package main

import (
	"testing"
)

func TestNewBankAccount_ValidBalance(t *testing.T) {
	account, err := NewBankAccount("12345", 100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if account.GetBalance() != 100.0 {
		t.Errorf("expected balance to be 100.0, got %f", account.GetBalance())
	}
}

func TestNewBankAccount_InvalidBalance(t *testing.T) {
	_, err := NewBankAccount("12345", -10.0)
	if err == nil {
		t.Fatalf("expected error for negative balance, got nil")
	}
}

func TestDeposit_ValidAmount(t *testing.T) {
	account, _ := NewBankAccount("12345", 100.0)
	err := account.Deposit(50.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if account.GetBalance() != 150.0 {
		t.Errorf("expected balance to be 150.0, got %f", account.GetBalance())
	}
}

func TestDeposit_InvalidAmount(t *testing.T) {
	account, _ := NewBankAccount("12345", 100.0)
	err := account.Deposit(0)
	if err == nil {
		t.Errorf("expected error for deposit of zero, got nil")
	}
}

func TestWithdraw_ValidAmount(t *testing.T) {
	account, _ := NewBankAccount("12345", 100.0)
	err := account.Withdraw(50.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if account.GetBalance() != 50.0 {
		t.Errorf("expected balance to be 50.0, got %f", account.GetBalance())
	}
}

func TestWithdraw_ExceedingBalance(t *testing.T) {
	account, _ := NewBankAccount("12345", 100.0)
	err := account.Withdraw(200.0)
	if err == nil || err.Error() != "insufficient funds" {
		t.Errorf("expected insufficient funds error, got %v", err)
	}
}

func TestWithdraw_InvalidAmount(t *testing.T) {
	account, _ := NewBankAccount("12345", 100.0)
	err := account.Withdraw(0)
	if err == nil {
		t.Errorf("expected error for withdrawal of zero, got nil")
	}
}
