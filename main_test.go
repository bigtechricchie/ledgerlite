package main

import "testing"

func TestCalculateBalance(t *testing.T) {
	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: "FUND-001",
			Type:      Deposit,
			Amount:    100000,
		},
		{
			ID:        "TXN-002",
			AccountID: "FUND-001",
			Type:      Withdrawal,
			Amount:    25000,
		},
		{
			ID:        "TXN-003",
			AccountID: "FUND-001",
			Type:      Deposit,
			Amount:    50000,
		},
	}

	balance := calculateBalance(transactions)

	var expected int64 = 125000

	if balance != expected {
		t.Errorf(
			"expected balance %d, got %d",
			expected,
			balance,
		)
	}
}

func TestValidateTransactionAcceptsValidTransaction(t *testing.T) {
	transaction := Transaction{
		ID:          "TXN-001",
		AccountID:   "FUND-001",
		Type:        Deposit,
		Amount:      10000,
		Description: "Funding",
	}

	err := validateTransaction(transaction)

	if err != nil {
		t.Errorf(
			"expected valid transaction, got error: %v",
			err,
		)
	}
}

func TestValidateTransactionRejectsNonPositiveAmount(t *testing.T) {
	transaction := Transaction{
		ID:          "TXN-001",
		AccountID:   "FUND-001",
		Type:        Deposit,
		Amount:      0,
		Description: "Funding",
	}

	err := validateTransaction(transaction)

	if err == nil {
		t.Fatal("expected transaction validation to fail")
	}

	expected := "transaction amount must be greater than zero"

	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}
}

func TestValidateLedgerRejectsDuplicateTransactionIDs(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: account.ID,
			Type:      Deposit,
			Amount:    100000,
		},
		{
			ID:        "TXN-001",
			AccountID: account.ID,
			Type:      Deposit,
			Amount:    50000,
		},
	}

	err := validateLedger(account, transactions)

	if err == nil {
		t.Fatal("expected ledger validation to fail")
	}

	expected := "duplicate transaction ID"

	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}
}

func TestValidateLedgerRejectsTransactionFromDifferentAccount(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: "FUND-002",
			Type:      Deposit,
			Amount:    100000,
		},
	}

	err := validateLedger(account, transactions)

	if err == nil {
		t.Fatal("expected ledger validation to fail")
	}

	expected := "transaction belongs to a different account"

	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}
}

func TestValidateLedgerRejectsWithdrawalExceedingBalance(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: account.ID,
			Type:      Deposit,
			Amount:    100000,
		},
		{
			ID:        "TXN-002",
			AccountID: account.ID,
			Type:      Withdrawal,
			Amount:    100001,
		},
	}

	err := validateLedger(account, transactions)

	if err == nil {
		t.Fatal("expected ledger validation to fail")
	}

	expected := "withdrawal exceeds available balance"

	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}
}

func TestValidateLedgerAllowsWithdrawalOfEntireBalance(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: account.ID,
			Type:      Deposit,
			Amount:    100000,
		},
		{
			ID:        "TXN-002",
			AccountID: account.ID,
			Type:      Withdrawal,
			Amount:    100000,
		},
	}

	err := validateLedger(account, transactions)

	if err != nil {
		t.Errorf(
			"expected ledger to be valid, got error: %v",
			err,
		)
	}
}

func TestAddTransactionAddsValidTransaction(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: account.ID,
			Type:      Deposit,
			Amount:    100000,
		},
	}

	transaction := Transaction{
		ID:          "TXN-002",
		AccountID:   account.ID,
		Type:        Deposit,
		Amount:      50000,
		Description: "Additional funding",
	}

	updatedTransactions, err := addTransaction(
		account,
		transactions,
		transaction,
	)

	if err != nil {
		t.Fatalf(
			"expected transaction to be added, got error: %v",
			err,
		)
	}

	expectedCount := 2

	if len(updatedTransactions) != expectedCount {
		t.Errorf(
			"expected %d transactions, got %d",
			expectedCount,
			len(updatedTransactions),
		)
	}
}

func TestAddTransactionRejectsDuplicateID(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: account.ID,
			Type:      Deposit,
			Amount:    100000,
		},
	}

	transaction := Transaction{
		ID:          "TXN-001",
		AccountID:   account.ID,
		Type:        Deposit,
		Amount:      50000,
		Description: "Duplicate",
	}

	updatedTransactions, err := addTransaction(
		account,
		transactions,
		transaction,
	)

	if err == nil {
		t.Fatal("expected transaction to be rejected")
	}

	if len(updatedTransactions) != len(transactions) {
		t.Errorf(
			"expected transaction count to remain %d, got %d",
			len(transactions),
			len(updatedTransactions),
		)
	}
}

func TestAddTransactionRejectsWithdrawalExceedingBalance(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: account.ID,
			Type:      Deposit,
			Amount:    100000,
		},
	}

	transaction := Transaction{
		ID:          "TXN-002",
		AccountID:   account.ID,
		Type:        Withdrawal,
		Amount:      100001,
		Description: "Excessive withdrawal",
	}

	updatedTransactions, err := addTransaction(
		account,
		transactions,
		transaction,
	)

	if err == nil {
		t.Fatal("expected transaction to be rejected")
	}

	expected := "withdrawal exceeds available balance"

	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}

	if len(updatedTransactions) != len(transactions) {
		t.Errorf(
			"expected transaction count to remain %d, got %d",
			len(transactions),
			len(updatedTransactions),
		)
	}
}
