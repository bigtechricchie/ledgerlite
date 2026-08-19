package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	ID          string
	Account     string
	Type        string
	Amount      float64
	Description string
}

func main() {
	transactions := []Transaction{
		{
			ID:          "TXN-001",
			Account:     "FUND-001",
			Type:        "DEPOSIT",
			Amount:      -10000,
			Description: "Initial funding",
		},
		{
			ID:          "TXN-002",
			Account:     "FUND-001",
			Type:        "WITHDRAWAL",
			Amount:      2500,
			Description: "Management fee",
		},
		{
			ID:          "TXN-003",
			Account:     "FUND-001",
			Type:        "DEPOSIT",
			Amount:      5000,
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

	balance := calculateBalance(transactions)

	fmt.Println("LedgerLite")
	fmt.Println("Simple Banking Ledger")
	fmt.Println()

	fmt.Println("Transactions")
	fmt.Println("------------")
	fmt.Println("Number of transactions:", len(transactions))

	for _, transaction := range transactions {
		fmt.Println(
			transaction.ID,
			transaction.Account,
			transaction.Type,
			transaction.Amount,
			transaction.Description,
		)
	}

	fmt.Println()
	fmt.Println("Balance:", balance)
}

func validateTransaction(transaction Transaction) error {
	if transaction.ID == "" {
		return errors.New("transaction ID is required")
	}

	if transaction.Account == "" {
		return errors.New("account is required")
	}

	if transaction.Type != "DEPOSIT" && transaction.Type != "WITHDRAWAL" {
		return errors.New("transaction type must be DEPOSIT or WITHDRAWAL")
	}

	if transaction.Amount <= 0 {
		return errors.New("transaction amount must be greater than zero")
	}

	return nil
}

func calculateBalance(transactions []Transaction) float64 {
	var balance float64

	for _, transaction := range transactions {
		if transaction.Type == "DEPOSIT" {
			balance += transaction.Amount
		}

		if transaction.Type == "WITHDRAWAL" {
			balance -= transaction.Amount
		}
	}

	return balance
}
