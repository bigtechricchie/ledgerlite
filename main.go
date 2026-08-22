package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	reader := bufio.NewReader(os.Stdin)

	runMenu(reader, account, transactions)
}

func runMenu(reader *bufio.Reader, account Account, transactions []Transaction) {
	shouldContinue := true

	for shouldContinue {
		printMenu()

		option, err := readLine(reader)
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
			transaction, err := createDeposit(reader, account)
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

		case "6":
			transaction, err := createWithdrawal(reader, account)
			if err != nil {
				fmt.Println("Could not record withdrawal")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			updatedTransactions := append(transactions, transaction)

			err = validateLedger(account, updatedTransactions)
			if err != nil {
				fmt.Println("Could not record withdrawal")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			transactions = updatedTransactions

			fmt.Println("Withdrawal recorded.")
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
	fmt.Println("[6] Record withdrawal")
	fmt.Println("[q] Quit")
	fmt.Println()
	fmt.Print("Select an option: ")
}

func createDeposit(reader *bufio.Reader, account Account) (Transaction, error) {
	fmt.Print("Transaction ID: ")

	transactionID, err := readLine(reader)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction ID")
	}

	fmt.Print("Amount in pence: ")

	rawAmount, err := readLine(reader)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction amount")
	}

	amount, err := strconv.ParseInt(rawAmount, 10, 64)
	if err != nil {
		return Transaction{}, errors.New("transaction amount must be a whole number of pence")
	}

	fmt.Print("Description: ")

	description, err := readLine(reader)
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

func createWithdrawal(reader *bufio.Reader, account Account) (Transaction, error) {
	fmt.Print("Transaction ID: ")

	transactionID, err := readLine(reader)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction ID")
	}

	fmt.Print("Amount in pence: ")

	rawAmount, err := readLine(reader)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction amount")
	}

	amount, err := strconv.ParseInt(rawAmount, 10, 64)
	if err != nil {
		return Transaction{}, errors.New("transaction amount must be a whole number of pence")
	}

	fmt.Print("Description: ")

	description, err := readLine(reader)
	if err != nil {
		return Transaction{}, errors.New("could not read transaction description")
	}

	transaction := Transaction{
		ID:          transactionID,
		AccountID:   account.ID,
		Type:        Withdrawal,
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

func readLine(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')

	if err != nil && err != io.EOF {
		return "", err
	}

	input = strings.TrimSpace(input)

	if err == io.EOF && input == "" {
		return "", io.EOF
	}

	return input, nil
}
