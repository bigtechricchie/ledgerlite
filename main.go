package main

import "fmt"

type Transaction struct {
	ID          string
	Account     string
	Type        string
	Amount      float64
	Description string
}

func main() {
	transaction := Transaction{
		ID:          "TXN-001",
		Account:     "FUND-001",
		Type:        "DEPOSIT",
		Amount:      10000,
		Description: "Initial funding",
	}

	fmt.Println("LedgerLite")
	fmt.Println("Simple Banking Ledger")
	fmt.Println()

	fmt.Println("Transaction")
	fmt.Println("-----------")
	fmt.Println("ID:", transaction.ID)
	fmt.Println("Account:", transaction.Account)
	fmt.Println("Type:", transaction.Type)
	fmt.Println("Amount:", transaction.Amount)
	fmt.Println("Description:", transaction.Description)
}
