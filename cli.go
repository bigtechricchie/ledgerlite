package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

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
			transaction, err := createTransaction(
				reader,
				account,
				transactions,
				Deposit,
			)
			if err != nil {
				fmt.Println("Could not record deposit")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			transactions, err = addTransaction(
				account,
				transactions,
				transaction,
			)
			if err != nil {
				fmt.Println("Could not record deposit")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			fmt.Println(
				"Deposit recorded:",
				transaction.ID,
			)
			fmt.Println()

		case "6":
			transaction, err := createTransaction(
				reader,
				account,
				transactions,
				Withdrawal,
			)
			if err != nil {
				fmt.Println("Could not record withdrawal")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			transactions, err = addTransaction(
				account,
				transactions,
				transaction,
			)
			if err != nil {
				fmt.Println("Could not record withdrawal")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			fmt.Println(
				"Withdrawal recorded:",
				transaction.ID,
			)
			fmt.Println()

		case "7":
			fmt.Print("Transaction ID: ")

			transactionID, err := readLine(reader)
			if err != nil {
				fmt.Println()
				fmt.Println("Could not read transaction ID")
				fmt.Println()
				continue
			}

			transaction, err := findTransactionByID(
				transactions,
				transactionID,
			)
			if err != nil {
				fmt.Println("Could not find transaction")
				fmt.Println("Reason:", err)
				fmt.Println()
				continue
			}

			fmt.Println()
			fmt.Println("Transaction")
			fmt.Println("-----------")
			printTransaction(transaction)
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
	fmt.Println("LedgerLite CLI")
	fmt.Println("Simple Banking Ledger")
	fmt.Println()
	fmt.Println("[1] View account")
	fmt.Println("[2] View transactions")
	fmt.Println("[3] View balance")
	fmt.Println("[4] View full report")
	fmt.Println("[5] Record deposit")
	fmt.Println("[6] Record withdrawal")
	fmt.Println("[7] Find transaction by ID")
	fmt.Println("[q] Quit")
	fmt.Println()
	fmt.Print("Select an option: ")
}

func createTransaction(
	reader *bufio.Reader,
	account Account,
	transactions []Transaction,
	transactionType TransactionType,
) (Transaction, error) {
	transactionID := nextTransactionID(transactions)

	fmt.Print("Amount (£): ")

	rawAmount, err := readLine(reader)
	if err != nil {
		return Transaction{}, errors.New(
			"could not read transaction amount",
		)
	}

	amount, err := parseAmount(rawAmount)
	if err != nil {
		return Transaction{}, err
	}

	fmt.Print("Description: ")

	description, err := readLine(reader)
	if err != nil {
		return Transaction{}, errors.New(
			"could not read transaction description",
		)
	}

	transaction := Transaction{
		ID:          transactionID,
		AccountID:   account.ID,
		Type:        transactionType,
		Amount:      amount,
		Description: description,
	}

	err = validateTransaction(transaction)
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func readLine(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')

	if err != nil && err != io.EOF {
		return "", err
	}

	input = strings.TrimSpace(input)

	// Preserve the final line when EOF arrives after readable input.
	if err == io.EOF && input == "" {
		return "", io.EOF
	}

	return input, nil
}

func printReport(account Account, transactions []Transaction, balance int64) {
	fmt.Println("LedgerLite CLI")
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
		printTransaction(transaction)
	}

	fmt.Println()
}

func printTransaction(transaction Transaction) {
	fmt.Println(
		transaction.ID,
		transaction.AccountID,
		transaction.Type,
		formatAmount(transaction.Amount),
		transaction.Description,
	)
}

func printBalance(balance int64) {
	fmt.Println("Balance")
	fmt.Println("-------")
	fmt.Println(formatAmount(balance))
	fmt.Println()
}
