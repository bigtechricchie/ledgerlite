package main

import (
	"errors"
	"strings"
)

func validateAccount(account Account) error {
	if strings.TrimSpace(account.ID) == "" {
		return errors.New("account ID is required")
	}

	if strings.TrimSpace(account.Name) == "" {
		return errors.New("account name is required")
	}

	return nil
}

func validateTransaction(transaction Transaction) error {
	if strings.TrimSpace(transaction.ID) == "" {
		return errors.New("transaction ID is required")
	}

	if strings.TrimSpace(transaction.AccountID) == "" {
		return errors.New("account ID is required")
	}

	if transaction.Type != Deposit && transaction.Type != Withdrawal {
		return errors.New("transaction type must be DEPOSIT or WITHDRAWAL")
	}

	if transaction.Amount <= 0 {
		return errors.New("transaction amount must be greater than zero")
	}

	if strings.TrimSpace(transaction.Description) == "" {
		return errors.New("transaction description is required")
	}

	return nil
}

func validateLedger(account Account, transactions []Transaction) error {
	// Transactions are validated in ledger order because withdrawals
	// depend on the balance available at that point in the sequence.
	err := validateAccount(account)
	if err != nil {
		return err
	}

	var balance int64
	seenTransactionIDs := map[string]bool{}

	for _, transaction := range transactions {
		err := validateTransaction(transaction)
		if err != nil {
			return err
		}

		if transaction.AccountID != account.ID {
			return errors.New("transaction belongs to a different account")
		}

		if seenTransactionIDs[transaction.ID] {
			return errors.New("duplicate transaction ID")
		}

		seenTransactionIDs[transaction.ID] = true

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
