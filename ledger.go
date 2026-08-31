package main

import (
	"errors"
	"fmt"
)

func addTransaction(
	account Account,
	transactions []Transaction,
	transaction Transaction,
) ([]Transaction, error) {
	err := validateTransaction(transaction)
	if err != nil {
		return transactions, err
	}

	updatedTransactions := append(transactions, transaction)

	err = validateLedger(account, updatedTransactions)
	if err != nil {
		return transactions, err
	}

	return updatedTransactions, nil
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

func findTransactionByID(
	transactions []Transaction,
	transactionID string,
) (Transaction, error) {
	for _, transaction := range transactions {
		if transaction.ID == transactionID {
			return transaction, nil
		}
	}

	return Transaction{}, errors.New("transaction not found")
}

func nextTransactionID(transactions []Transaction) string {
	usedIDs := map[string]bool{}

	for _, transaction := range transactions {
		usedIDs[transaction.ID] = true
	}

	for number := 1; ; number++ {
		transactionID := fmt.Sprintf("TXN-%03d", number)

		if !usedIDs[transactionID] {
			return transactionID
		}
	}
}
