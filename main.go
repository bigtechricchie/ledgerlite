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

	runMenu(account, transactions)
}

func runMenu(account Account, transactions []Transaction) {
	shouldContinue := true

	for shouldContinue {
		printMenu()

		var option string

		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println()
			fmt.Println("Input stream closed.")
			return
		}

		fmt.Println()

		switch option {
		case "1":
			printAccount(account)

		case "2":
			printTransactions(transactions)

		case "3":
			balance := calculateBalance(transactions)
			printBalance(balance)

		case "4":
			balance := calculateBalance(transactions)
			printReport(account, transactions, balance)

		case "5":
			transaction, err := createDeposit(account)
			if err != nil {
				fmt.Println("Could not record deposit")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			updatedTransactions := append(transactions, transaction)

			err = validateLedger(account, updatedTransactions)
			if err != nil {
				fmt.Println("Could not record deposit")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			transactions = updatedTransactions

			fmt.Println("Deposit recorded.")
			fmt.Println()

		case "q":
			shouldContinue = false

		default:
			fmt.Println("Unknown option")
			fmt.Println()
		}
	}
}

func printMenu() {
	fmt.Println("LedgerLite")
	fmt.Println("Simple Banking Ledger")
	fmt.Println()
	fmt.Println("[1] View account")
	fmt.Println("[2] View transactions")
	fmt.Println("[3] View balance")
	fmt.Println("[4] View full report")
	fmt.Println("[5] Record deposit")
	fmt.Println("[q] Quit")
	fmt.Println()
	fmt.Print("Select an option: ")
}

func createDeposit(account Account) (Transaction, error) {
	var transactionID string
	var amount int64
	var description string

	fmt.Print("Transaction ID: ")
	_, err := fmt.Scanln(&transactionID)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction ID")
	}

	fmt.Print("Amount in pence: ")
	_, err = fmt.Scanln(&amount)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction amount")
	}

	fmt.Print("Description: ")
	_, err = fmt.Scanln(&description)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction description")
	}

	transaction := Transaction{
		ID:          transactionID,
		AccountID:   account.ID,
		Type:        Deposit,
		Amount:      amount,
		Description: description,
	}

	err = validateTransaction(transaction)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
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
	seenTransactionIDs := map[string]bool{}

	for _, transaction := range transactions {
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

func printReport(account Account, transactions []Transaction, balance int64) {
	fmt.Println("LedgerLite")
	fmt.Println("Simple Banking Ledger")
	fmt.Println()

	printAccount(account)
	printTransactions(transactions)
	printBalance(balance)
}

func printAccount(account Account) {
	fmt.Println("Account")
	fmt.Println("-------")
	fmt.Println("ID:", account.ID)
	fmt.Println("Name:", account.Name)
	fmt.Println()
}

func printTransactions(transactions []Transaction) {
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
}

func printBalance(balance int64) {
	fmt.Println("Balance")
	fmt.Println("-------")
	fmt.Println(formatAmount(balance))
	fmt.Println()
}

func formatAmount(amount int64) string {
	pounds := amount / 100
	pence := amount % 100

	return fmt.Sprintf("£%d.%02d", pounds, pence)
}
