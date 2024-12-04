package main

import (
	"errors"
)

type BankAccount struct {
	accountNumber string
	balance       float64
}

// Конструктор
func NewBankAccount(accountNumber string, balance float64) (*BankAccount, error) {
	if balance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}
	return &BankAccount{accountNumber: accountNumber, balance: balance}, nil
}

// Метод для пополнения баланса
func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	ba.balance += amount
	return nil
}

// Метод для снятия средств
func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > ba.balance {
		return errors.New("insufficient funds")
	}
	ba.balance -= amount
	return nil
}

// Метод для получения текущего баланса
func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}
