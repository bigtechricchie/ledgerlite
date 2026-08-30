package main

import "testing"

func TestValidateAccount(t *testing.T) {

	tests := []struct {
		name          string
		account       Account
		expectedError string
	}{
		{
			name: "valid account",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			expectedError: "",
		},
		{
			name: "missing account ID",
			account: Account{
				ID:   "",
				Name: "LedgerLite Demo Fund",
			},
			expectedError: "account ID is required",
		},
		{
			name: "whitespace account ID",
			account: Account{
				ID:   "   ",
				Name: "LedgerLite Demo Fund",
			},
			expectedError: "account ID is required",
		},
		{
			name: "whitespace account name",
			account: Account{
				ID:   "FUND-001",
				Name: "   ",
			},
			expectedError: "account name is required",
		},
		{
			name: "missing account name",
			account: Account{
				ID:   "FUND-001",
				Name: "",
			},
			expectedError: "account name is required",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateAccount(test.account)

			if test.expectedError == "" {
				if err != nil {
					t.Errorf(
						"expected account to be valid, got error: %v",
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
			name: "whitespace transaction ID",
			transaction: Transaction{
				ID:          "   ",
				AccountID:   "FUND-001",
				Type:        Deposit,
				Amount:      10000,
				Description: "Investor capital subscription",
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
			name: "whitespace account ID",
			transaction: Transaction{
				ID:          "TXN-001",
				AccountID:   "   ",
				Type:        Deposit,
				Amount:      10000,
				Description: "Investor capital subscription",
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
		{
			name: "missing description",
			transaction: Transaction{
				ID:          "TXN-001",
				AccountID:   "FUND-001",
				Type:        Deposit,
				Amount:      10000,
				Description: "",
			},
			expectedError: "transaction description is required",
		},
		{
			name: "whitespace description",
			transaction: Transaction{
				ID:          "TXN-001",
				AccountID:   "FUND-001",
				Type:        Deposit,
				Amount:      10000,
				Description: "   ",
			},
			expectedError: "transaction description is required",
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
	tests := []struct {
		name          string
		transactions  []Transaction
		account       Account
		expectedError string
	}{
		{
			name: "valid ledger",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      100000,
					Description: "Initial funding",
				},
				{
					ID:          "TXN-002",
					AccountID:   "FUND-001",
					Type:        Withdrawal,
					Amount:      25000,
					Description: "Management fee",
				},
			},
			expectedError: "",
		},
		{
			name: "duplicate transaction ID",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      100000,
					Description: "Initial funding",
				},
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      50000,
					Description: "Additional funding",
				},
			},
			expectedError: "duplicate transaction ID",
		},
		{
			name: "transaction belongs to different account",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-002",
					Type:        Deposit,
					Amount:      100000,
					Description: "Initial funding",
				},
			},
			expectedError: "transaction belongs to a different account",
		},
		{
			name: "invalid transaction amount",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      0,
					Description: "Initial funding",
				},
			},
			expectedError: "transaction amount must be greater than zero",
		},
		{
			name: "withdrawal exceeds balance",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      100000,
					Description: "Initial funding",
				},
				{
					ID:          "TXN-002",
					AccountID:   "FUND-001",
					Type:        Withdrawal,
					Amount:      100001,
					Description: "Excessive withdrawal",
				},
			},
			expectedError: "withdrawal exceeds available balance",
		},
		{
			name: "withdrawal equals full balance",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      100000,
					Description: "Initial funding",
				},
				{
					ID:          "TXN-002",
					AccountID:   "FUND-001",
					Type:        Withdrawal,
					Amount:      100000,
					Description: "Full redemption",
				},
			},
			expectedError: "",
		},
		{
			name: "missing account ID",
			account: Account{
				ID:   "",
				Name: "LedgerLite Demo Fund",
			},
			transactions:  []Transaction{},
			expectedError: "account ID is required",
		},
		{
			name: "missing account name",
			account: Account{
				ID:   "FUND-001",
				Name: "",
			},
			transactions:  []Transaction{},
			expectedError: "account name is required",
		},
		{
			name: "transaction missing description",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      100000,
					Description: "",
				},
			},
			expectedError: "transaction description is required",
		},
		{
			name: "withdrawal before funding",
			account: Account{
				ID:   "FUND-001",
				Name: "LedgerLite Demo Fund",
			},
			transactions: []Transaction{
				{
					ID:          "TXN-001",
					AccountID:   "FUND-001",
					Type:        Withdrawal,
					Amount:      10000,
					Description: "Early withdrawal",
				},
				{
					ID:          "TXN-002",
					AccountID:   "FUND-001",
					Type:        Deposit,
					Amount:      10000,
					Description: "Later funding",
				},
			},
			expectedError: "withdrawal exceeds available balance",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateLedger(test.account, test.transactions)

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
		Name: "LedgerLite Demo Fund",
	}

	transactions := []Transaction{
		{
			ID:          "TXN-001",
			AccountID:   account.ID,
			Type:        Deposit,
			Amount:      100000,
			Description: "Additional funding",
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
		Name: "LedgerLite Demo Fund",
	}

	transactions := []Transaction{
		{
			ID:          "TXN-001",
			AccountID:   account.ID,
			Type:        Deposit,
			Amount:      100000,
			Description: "Initial funding",
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
		Name: "LedgerLite Demo Fund",
	}

	transactions := []Transaction{
		{
			ID:          "TXN-001",
			AccountID:   account.ID,
			Type:        Deposit,
			Amount:      100000,
			Description: "Initial funding",
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

func TestFindTransactionByID(t *testing.T) {
	transactions := []Transaction{
		{
			ID:        "TXN-001",
			AccountID: "FUND-001",
			Type:      Deposit,
			Amount:    100000,
		},
		{
			ID:          "TXN-002",
			AccountID:   "FUND-001",
			Type:        Withdrawal,
			Amount:      25000,
			Description: "Management fee",
		},
	}

	tests := []struct {
		name          string
		transactionID string
		expectedID    string
		expectedError string
	}{
		{
			name:          "existing transaction",
			transactionID: "TXN-002",
			expectedID:    "TXN-002",
			expectedError: "",
		},
		{
			name:          "missing transaction",
			transactionID: "TXN-999",
			expectedID:    "",
			expectedError: "transaction not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			transaction, err := findTransactionByID(
				transactions,
				test.transactionID,
			)

			if test.expectedError != "" {
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

				return
			}

			if err != nil {
				t.Fatalf(
					"expected transaction to be found, got error: %v",
					err,
				)
			}

			if transaction.ID != test.expectedID {
				t.Errorf(
					"expected transaction ID %q, got %q",
					test.expectedID,
					transaction.ID,
				)
			}
		})
	}
}

func TestNextTransactionID(t *testing.T) {
	tests := []struct {
		name         string
		transactions []Transaction
		expectedID   string
	}{
		{
			name:         "empty ledger",
			transactions: []Transaction{},
			expectedID:   "TXN-001",
		},
		{
			name: "next sequential ID",
			transactions: []Transaction{
				{ID: "TXN-001"},
				{ID: "TXN-002"},
				{ID: "TXN-003"},
			},
			expectedID: "TXN-004",
		},
		{
			name: "fills first available gap",
			transactions: []Transaction{
				{ID: "TXN-001"},
				{ID: "TXN-003"},
			},
			expectedID: "TXN-002",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			transactionID := nextTransactionID(
				test.transactions,
			)

			if transactionID != test.expectedID {
				t.Errorf(
					"expected transaction ID %q, got %q",
					test.expectedID,
					transactionID,
				)
			}
		})
	}
}

func TestParseAmount(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedAmount int64
		expectedError  string
	}{
		{
			name:           "whole pounds",
			input:          "1250",
			expectedAmount: 125000,
			expectedError:  "",
		},
		{
			name:           "pounds and pence",
			input:          "1250.50",
			expectedAmount: 125050,
			expectedError:  "",
		},
		{
			name:           "single decimal place",
			input:          "1250.5",
			expectedAmount: 125050,
			expectedError:  "",
		},
		{
			name:           "one penny",
			input:          "0.01",
			expectedAmount: 1,
			expectedError:  "",
		},
		{
			name:           "trailing decimal point",
			input:          "10.",
			expectedAmount: 1000,
			expectedError:  "",
		},
		{
			name:          "too many decimal places",
			input:         "10.001",
			expectedError: "transaction amount cannot have more than two decimal places",
		},
		{
			name:          "multiple decimal points",
			input:         "10.00.1",
			expectedError: "transaction amount must be a valid currency amount",
		},
		{
			name:          "non numeric input",
			input:         "hello",
			expectedError: "transaction amount must be a valid currency amount",
		},
		{
			name:          "currency symbol",
			input:         "£10.00",
			expectedError: "transaction amount must be a valid currency amount",
		},
		{
			name:          "negative amount",
			input:         "-10.50",
			expectedError: "transaction amount must be greater than zero",
		},
		{
			name:          "zero amount",
			input:         "0.00",
			expectedError: "transaction amount must be greater than zero",
		},
		{
			name:          "missing pounds",
			input:         ".50",
			expectedError: "transaction amount must include pounds",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			amount, err := parseAmount(test.input)

			if test.expectedError != "" {
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

				return
			}

			if err != nil {
				t.Fatalf(
					"expected amount to be valid, got error: %v",
					err,
				)
			}

			if amount != test.expectedAmount {
				t.Errorf(
					"expected amount %d, got %d",
					test.expectedAmount,
					amount,
				)
			}
		})
	}
}
