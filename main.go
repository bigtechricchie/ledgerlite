package main

import (
	"errors"
	"fmt"
)

type TransactionType string

const (
	Deposit    TransactionType = "DEPOSIT"
	Withdrawal TransactionType = "WITHDRAWAL"
)

type Account struct {
	ID   string
	Name string
}

type Transaction struct {
	ID          string
	AccountID   string
	Type        TransactionType
	Amount      int64
	Description string
}

func main() {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:          "TXN-001",
			AccountID:   account.ID,
			Type:        Deposit,
			Amount:      1000000,
			Description: "Initial funding",
		},
		{
			ID:          "TXN-002",
			AccountID:   account.ID,
			Type:        Withdrawal,
			Amount:      250000,
			Description: "Management fee",
		},
		{
			ID:          "TXN-003",
			AccountID:   account.ID,
			Type:        Deposit,
			Amount:      500000,
			Description: "Additional funding",
		},
	}

	for _, transaction := range transactions {
		err := validateTransaction(transaction)
		if err != nil {
			fmt.Println("Invalid transaction:", transaction.ID)
			fmt.Println("Reason:", err)
			return
		}
	}

	err := validateLedger(account, transactions)
	if err != nil {
		fmt.Println("Invalid ledger")
		fmt.Println("Reason:", err)
		return
	}

	balance := calculateBalance(transactions)

	fmt.Println("LedgerLite")
	fmt.Println("Simple Banking Ledger")
	fmt.Println()

	fmt.Println("Account")
	fmt.Println("-------")
	fmt.Println("ID:", account.ID)
	fmt.Println("Name:", account.Name)
	fmt.Println()

	fmt.Println("Transactions")
	fmt.Println("------------")
	fmt.Println("Number of transactions:", len(transactions))

	for _, transaction := range transactions {
		fmt.Println(
			transaction.ID,
			transaction.AccountID,
			transaction.Type,
			formatAmount(transaction.Amount),
			transaction.Description,
		)
	}

	fmt.Println()
	fmt.Println("Balance:", formatAmount(balance))
}

func validateTransaction(transaction Transaction) error {
	if transaction.ID == "" {
		return errors.New("transaction ID is required")
	}

	if transaction.AccountID == "" {
		return errors.New("account ID is required")
	}

	if transaction.Type != Deposit && transaction.Type != Withdrawal {
		return errors.New("transaction type must be DEPOSIT or WITHDRAWAL")
	}

	if transaction.Amount <= 0 {
		return errors.New("transaction amount must be greater than zero")
	}

	return nil
}

func validateLedger(account Account, transactions []Transaction) error {
	var balance int64

	for _, transaction := range transactions {
		if transaction.AccountID != account.ID {
			return errors.New("transaction belongs to a different account")
		}

		if transaction.Type == Deposit {
			balance += transaction.Amount
		}

		if transaction.Type == Withdrawal {
			if transaction.Amount > balance {
				return errors.New("withdrawal exceeds available balance")
			}

			balance -= transaction.Amount
		}
	}

	return nil
}

func calculateBalance(transactions []Transaction) int64 {
	var balance int64

	for _, transaction := range transactions {
		if transaction.Type == Deposit {
			balance += transaction.Amount
		}

		if transaction.Type == Withdrawal {
			balance -= transaction.Amount
		}
	}

	return balance
}

func formatAmount(amount int64) string {
	pounds := amount / 100
	pence := amount % 100

	return fmt.Sprintf("£%d.%02d", pounds, pence)
}
