package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	account := Account{
		ID:   "FUND-001",
		Name: "LedgerLite Demo Fund",
	}

	err := validateAccount(account)
	if err != nil {
		fmt.Println("Invalid account")
		fmt.Println("Reason:", err)
		return
	}

	transactions := []Transaction{
		{
			ID:          "TXN-001",
			AccountID:   account.ID,
			Type:        Deposit,
			Amount:      250000000,
			Description: "Investor capital subscription",
		},
		{
			ID:          "TXN-002",
			AccountID:   account.ID,
			Type:        Withdrawal,
			Amount:      1850000,
			Description: "Quarterly management fee",
		},
		{
			ID:          "TXN-003",
			AccountID:   account.ID,
			Type:        Withdrawal,
			Amount:      725000,
			Description: "Fund administration expense",
		},
		{
			ID:          "TXN-004",
			AccountID:   account.ID,
			Type:        Deposit,
			Amount:      8750000,
			Description: "Interest income settlement",
		},
		{
			ID:          "TXN-005",
			AccountID:   account.ID,
			Type:        Withdrawal,
			Amount:      3200000,
			Description: "Investor redemption payment",
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

	err = validateLedger(account, transactions)
	if err != nil {
		fmt.Println("Invalid ledger")
		fmt.Println("Reason:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	runMenu(reader, account, transactions)
}
