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

func TestValidateTransaction(t *testing.T) {
	tests := []struct {
		name          string
		transaction   Transaction
		expectedError string
	}{
		{
			name: "valid deposit",
			transaction: Transaction{
				ID:          "TXN-001",
				AccountID:   "FUND-001",
				Type:        Deposit,
				Amount:      10000,
				Description: "Funding",
			},
			expectedError: "",
		},
		{
			name: "missing transaction ID",
			transaction: Transaction{
				AccountID:   "FUND-001",
				Type:        Deposit,
				Amount:      10000,
				Description: "Funding",
			},
			expectedError: "transaction ID is required",
		},
		{
			name: "missing account ID",
			transaction: Transaction{
				ID:          "TXN-001",
				Type:        Deposit,
				Amount:      10000,
				Description: "Funding",
			},
			expectedError: "account ID is required",
		},
		{
			name: "invalid transaction type",
			transaction: Transaction{
				ID:          "TXN-001",
				AccountID:   "FUND-001",
				Type:        TransactionType("TRANSFER"),
				Amount:      10000,
				Description: "Funding",
			},
			expectedError: "transaction type must be DEPOSIT or WITHDRAWAL",
		},
		{
			name: "zero amount",
			transaction: Transaction{
				ID:          "TXN-001",
				AccountID:   "FUND-001",
				Type:        Deposit,
				Amount:      0,
				Description: "Funding",
			},
			expectedError: "transaction amount must be greater than zero",
		},
		{
			name: "negative amount",
			transaction: Transaction{
				ID:          "TXN-001",
				AccountID:   "FUND-001",
				Type:        Withdrawal,
				Amount:      -1,
				Description: "Invalid withdrawal",
			},
			expectedError: "transaction amount must be greater than zero",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateTransaction(test.transaction)

			if test.expectedError == "" {
				if err != nil {
					t.Errorf(
						"expected transaction to be valid, got error: %v",
						err,
					)
				}

				return
			}

			if err == nil {
				t.Fatalf(
					"expected error %q, got nil",
					test.expectedError,
				)
			}

			if err.Error() != test.expectedError {
				t.Errorf(
					"expected error %q, got %q",
					test.expectedError,
					err.Error(),
				)
			}
		})
	}
}

func TestValidateLedger(t *testing.T) {
	account := Account{
		ID:   "FUND-001",
		Name: "Global Macro Fund",
	}

	tests := []struct {
		name          string
		transactions  []Transaction
		expectedError string
	}{
		{
			name: "valid ledger",
			transactions: []Transaction{
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
					Amount:    25000,
				},
			},
			expectedError: "",
		},
		{
			name: "duplicate transaction ID",
			transactions: []Transaction{
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
			},
			expectedError: "duplicate transaction ID",
		},
		{
			name: "transaction belongs to different account",
			transactions: []Transaction{
				{
					ID:        "TXN-001",
					AccountID: "FUND-002",
					Type:      Deposit,
					Amount:    100000,
				},
			},
			expectedError: "transaction belongs to a different account",
		},
		{
			name: "withdrawal exceeds balance",
			transactions: []Transaction{
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
			},
			expectedError: "withdrawal exceeds available balance",
		},
		{
			name: "withdrawal equals full balance",
			transactions: []Transaction{
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
			},
			expectedError: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateLedger(account, test.transactions)

			if test.expectedError == "" {
				if err != nil {
					t.Errorf(
						"expected ledger to be valid, got error: %v",
						err,
					)
				}

				return
			}

			if err == nil {
				t.Fatalf(
					"expected error %q, got nil",
					test.expectedError,
				)
			}

			if err.Error() != test.expectedError {
				t.Errorf(
					"expected error %q, got %q",
					test.expectedError,
					err.Error(),
				)
			}
		})
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
