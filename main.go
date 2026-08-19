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
	transactions := []Transaction{
		{
			ID:          "TXN-001",
			Account:     "FUND-001",
			Type:        "DEPOSIT",
			Amount:      10000,
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
}
